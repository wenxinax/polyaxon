# -*- coding: utf-8 -*-
from __future__ import absolute_import, division, print_function

import logging

from docker.errors import DockerException

from api.celery_api import app as celery_app
from api.settings import CeleryTasks
from repos import dockerize
from repos.models import Repo

from spawner import scheduler
from spawner.utils.constants import ExperimentLifeCycle
from experiments.models import Experiment

logger = logging.getLogger('polyaxon.tasks.experiments')


def get_valid_experiment(experiment_id):
    try:
        experiment = Experiment.objects.get(id=experiment_id)
    except Experiment.DoesNotExist:
        logger.info('Experiment id `{}` does not exist'.format(experiment_id))
        return None

    if experiment.is_done:
        logger.info('Experiment id `{}` stopped with status `{}`.'.format(experiment_id,
                                                                          experiment.last_status))
        return None

    return experiment


@celery_app.task(name=CeleryTasks.EXPERIMENTS_BUILD)
def build_experiment(experiment_id):
    experiment = get_valid_experiment(experiment_id=experiment_id)
    if not experiment:
        return

    # No need to build the image, start the experiment directly
    if not experiment.compiled_spec.run_exec:
        start_experiment.delay(experiment_id=experiment_id)
        return

    # Update experiment status to show that its building
    experiment.set_status(ExperimentLifeCycle.BUILDING)

    # docker image
    try:
        status = dockerize.build_experiment(experiment)
    except DockerException as e:
        logger.warning('Failed to build experiment %s\n' % e)
        experiment.set_status(ExperimentLifeCycle.FAILED)
        return
    except Repo.DoesNotExist:
        logger.warning('No code was found for this project')
        experiment.set_status(ExperimentLifeCycle.FAILED)
        return

    if not status:
        return

    # Now we can start the experiment
    start_experiment.delay(experiment_id=experiment_id)


@celery_app.task(name=CeleryTasks.EXPERIMENTS_START)
def start_experiment(experiment_id):
    experiment = get_valid_experiment(experiment_id=experiment_id)
    if not experiment:
        return

    scheduler.start_experiment(experiment)


@celery_app.task(name=CeleryTasks.EXPERIMENTS_STOP)
def stop_experiment(experiment_id):
    experiment = get_valid_experiment(experiment_id=experiment_id)
    if not experiment:
        return

    scheduler.stop_experiment(experiment, update_status=True)


@celery_app.task(name=CeleryTasks.EXPERIMENTS_CHECK_STATUS)
def check_experiment_status(experiment_uuid):
    experiment = Experiment.objects.get(uuid=experiment_uuid)
    experiment.update_status()
