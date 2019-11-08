#!/usr/bin/python
#
# Copyright 2019 Polyaxon, Inc.
#
# Licensed under the Apache License, Version 2.0 (the "License");
# you may not use this file except in compliance with the License.
# You may obtain a copy of the License at
#
#      http://www.apache.org/licenses/LICENSE-2.0
#
# Unless required by applicable law or agreed to in writing, software
# distributed under the License is distributed on an "AS IS" BASIS,
# WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
# See the License for the specific language governing permissions and
# limitations under the License.

# coding: utf-8

"""
    Polyaxon sdk

    No description provided (generated by Swagger Codegen https://github.com/swagger-api/swagger-codegen)  # noqa: E501

    OpenAPI spec version: 1.14.4
    Contact: contact@polyaxon.com
    Generated by: https://github.com/swagger-api/swagger-codegen.git
"""


import pprint
import re  # noqa: F401

import six


class V1RunMetaInfo(object):
    """NOTE: This class is auto generated by the swagger code generator program.

    Do not edit the class manually.
    """

    """
    Attributes:
      swagger_types (dict): The key is attribute name
                            and the value is attribute type.
      attribute_map (dict): The key is attribute name
                            and the value is json key in definition.
    """
    swagger_types = {
        'service': 'bool',
        'workflow_strategy': 'str',
        'workflow_concurrency': 'int'
    }

    attribute_map = {
        'service': 'service',
        'workflow_strategy': 'workflow_strategy',
        'workflow_concurrency': 'workflow_concurrency'
    }

    def __init__(self, service=None, workflow_strategy=None, workflow_concurrency=None):  # noqa: E501
        """V1RunMetaInfo - a model defined in Swagger"""  # noqa: E501

        self._service = None
        self._workflow_strategy = None
        self._workflow_concurrency = None
        self.discriminator = None

        if service is not None:
            self.service = service
        if workflow_strategy is not None:
            self.workflow_strategy = workflow_strategy
        if workflow_concurrency is not None:
            self.workflow_concurrency = workflow_concurrency

    @property
    def service(self):
        """Gets the service of this V1RunMetaInfo.  # noqa: E501


        :return: The service of this V1RunMetaInfo.  # noqa: E501
        :rtype: bool
        """
        return self._service

    @service.setter
    def service(self, service):
        """Sets the service of this V1RunMetaInfo.


        :param service: The service of this V1RunMetaInfo.  # noqa: E501
        :type: bool
        """

        self._service = service

    @property
    def workflow_strategy(self):
        """Gets the workflow_strategy of this V1RunMetaInfo.  # noqa: E501


        :return: The workflow_strategy of this V1RunMetaInfo.  # noqa: E501
        :rtype: str
        """
        return self._workflow_strategy

    @workflow_strategy.setter
    def workflow_strategy(self, workflow_strategy):
        """Sets the workflow_strategy of this V1RunMetaInfo.


        :param workflow_strategy: The workflow_strategy of this V1RunMetaInfo.  # noqa: E501
        :type: str
        """

        self._workflow_strategy = workflow_strategy

    @property
    def workflow_concurrency(self):
        """Gets the workflow_concurrency of this V1RunMetaInfo.  # noqa: E501


        :return: The workflow_concurrency of this V1RunMetaInfo.  # noqa: E501
        :rtype: int
        """
        return self._workflow_concurrency

    @workflow_concurrency.setter
    def workflow_concurrency(self, workflow_concurrency):
        """Sets the workflow_concurrency of this V1RunMetaInfo.


        :param workflow_concurrency: The workflow_concurrency of this V1RunMetaInfo.  # noqa: E501
        :type: int
        """

        self._workflow_concurrency = workflow_concurrency

    def to_dict(self):
        """Returns the model properties as a dict"""
        result = {}

        for attr, _ in six.iteritems(self.swagger_types):
            value = getattr(self, attr)
            if isinstance(value, list):
                result[attr] = list(map(
                    lambda x: x.to_dict() if hasattr(x, "to_dict") else x,
                    value
                ))
            elif hasattr(value, "to_dict"):
                result[attr] = value.to_dict()
            elif isinstance(value, dict):
                result[attr] = dict(map(
                    lambda item: (item[0], item[1].to_dict())
                    if hasattr(item[1], "to_dict") else item,
                    value.items()
                ))
            else:
                result[attr] = value
        if issubclass(V1RunMetaInfo, dict):
            for key, value in self.items():
                result[key] = value

        return result

    def to_str(self):
        """Returns the string representation of the model"""
        return pprint.pformat(self.to_dict())

    def __repr__(self):
        """For `print` and `pprint`"""
        return self.to_str()

    def __eq__(self, other):
        """Returns true if both objects are equal"""
        if not isinstance(other, V1RunMetaInfo):
            return False

        return self.__dict__ == other.__dict__

    def __ne__(self, other):
        """Returns true if both objects are not equal"""
        return not self == other