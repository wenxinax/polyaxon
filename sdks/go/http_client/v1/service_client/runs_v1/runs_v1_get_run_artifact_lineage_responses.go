// Copyright 2018-2020 Polyaxon, Inc.
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//      http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

// Code generated by go-swagger; DO NOT EDIT.

package runs_v1

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"
	"io"

	"github.com/go-openapi/runtime"
	"github.com/go-openapi/strfmt"

	"github.com/polyaxon/polyaxon/sdks/go/http_client/v1/service_model"
)

// RunsV1GetRunArtifactLineageReader is a Reader for the RunsV1GetRunArtifactLineage structure.
type RunsV1GetRunArtifactLineageReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *RunsV1GetRunArtifactLineageReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewRunsV1GetRunArtifactLineageOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 204:
		result := NewRunsV1GetRunArtifactLineageNoContent()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 403:
		result := NewRunsV1GetRunArtifactLineageForbidden()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 404:
		result := NewRunsV1GetRunArtifactLineageNotFound()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	default:
		result := NewRunsV1GetRunArtifactLineageDefault(response.Code())
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		if response.Code()/100 == 2 {
			return result, nil
		}
		return nil, result
	}
}

// NewRunsV1GetRunArtifactLineageOK creates a RunsV1GetRunArtifactLineageOK with default headers values
func NewRunsV1GetRunArtifactLineageOK() *RunsV1GetRunArtifactLineageOK {
	return &RunsV1GetRunArtifactLineageOK{}
}

/*RunsV1GetRunArtifactLineageOK handles this case with default header values.

A successful response.
*/
type RunsV1GetRunArtifactLineageOK struct {
	Payload *service_model.V1RunArtifact
}

func (o *RunsV1GetRunArtifactLineageOK) Error() string {
	return fmt.Sprintf("[GET /api/v1/{owner}/{project}/runs/{uuid}/artifacts_lineage/{name}][%d] runsV1GetRunArtifactLineageOK  %+v", 200, o.Payload)
}

func (o *RunsV1GetRunArtifactLineageOK) GetPayload() *service_model.V1RunArtifact {
	return o.Payload
}

func (o *RunsV1GetRunArtifactLineageOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(service_model.V1RunArtifact)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewRunsV1GetRunArtifactLineageNoContent creates a RunsV1GetRunArtifactLineageNoContent with default headers values
func NewRunsV1GetRunArtifactLineageNoContent() *RunsV1GetRunArtifactLineageNoContent {
	return &RunsV1GetRunArtifactLineageNoContent{}
}

/*RunsV1GetRunArtifactLineageNoContent handles this case with default header values.

No content.
*/
type RunsV1GetRunArtifactLineageNoContent struct {
	Payload interface{}
}

func (o *RunsV1GetRunArtifactLineageNoContent) Error() string {
	return fmt.Sprintf("[GET /api/v1/{owner}/{project}/runs/{uuid}/artifacts_lineage/{name}][%d] runsV1GetRunArtifactLineageNoContent  %+v", 204, o.Payload)
}

func (o *RunsV1GetRunArtifactLineageNoContent) GetPayload() interface{} {
	return o.Payload
}

func (o *RunsV1GetRunArtifactLineageNoContent) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// response payload
	if err := consumer.Consume(response.Body(), &o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewRunsV1GetRunArtifactLineageForbidden creates a RunsV1GetRunArtifactLineageForbidden with default headers values
func NewRunsV1GetRunArtifactLineageForbidden() *RunsV1GetRunArtifactLineageForbidden {
	return &RunsV1GetRunArtifactLineageForbidden{}
}

/*RunsV1GetRunArtifactLineageForbidden handles this case with default header values.

You don't have permission to access the resource.
*/
type RunsV1GetRunArtifactLineageForbidden struct {
	Payload interface{}
}

func (o *RunsV1GetRunArtifactLineageForbidden) Error() string {
	return fmt.Sprintf("[GET /api/v1/{owner}/{project}/runs/{uuid}/artifacts_lineage/{name}][%d] runsV1GetRunArtifactLineageForbidden  %+v", 403, o.Payload)
}

func (o *RunsV1GetRunArtifactLineageForbidden) GetPayload() interface{} {
	return o.Payload
}

func (o *RunsV1GetRunArtifactLineageForbidden) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// response payload
	if err := consumer.Consume(response.Body(), &o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewRunsV1GetRunArtifactLineageNotFound creates a RunsV1GetRunArtifactLineageNotFound with default headers values
func NewRunsV1GetRunArtifactLineageNotFound() *RunsV1GetRunArtifactLineageNotFound {
	return &RunsV1GetRunArtifactLineageNotFound{}
}

/*RunsV1GetRunArtifactLineageNotFound handles this case with default header values.

Resource does not exist.
*/
type RunsV1GetRunArtifactLineageNotFound struct {
	Payload interface{}
}

func (o *RunsV1GetRunArtifactLineageNotFound) Error() string {
	return fmt.Sprintf("[GET /api/v1/{owner}/{project}/runs/{uuid}/artifacts_lineage/{name}][%d] runsV1GetRunArtifactLineageNotFound  %+v", 404, o.Payload)
}

func (o *RunsV1GetRunArtifactLineageNotFound) GetPayload() interface{} {
	return o.Payload
}

func (o *RunsV1GetRunArtifactLineageNotFound) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// response payload
	if err := consumer.Consume(response.Body(), &o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewRunsV1GetRunArtifactLineageDefault creates a RunsV1GetRunArtifactLineageDefault with default headers values
func NewRunsV1GetRunArtifactLineageDefault(code int) *RunsV1GetRunArtifactLineageDefault {
	return &RunsV1GetRunArtifactLineageDefault{
		_statusCode: code,
	}
}

/*RunsV1GetRunArtifactLineageDefault handles this case with default header values.

An unexpected error response
*/
type RunsV1GetRunArtifactLineageDefault struct {
	_statusCode int

	Payload *service_model.RuntimeError
}

// Code gets the status code for the runs v1 get run artifact lineage default response
func (o *RunsV1GetRunArtifactLineageDefault) Code() int {
	return o._statusCode
}

func (o *RunsV1GetRunArtifactLineageDefault) Error() string {
	return fmt.Sprintf("[GET /api/v1/{owner}/{project}/runs/{uuid}/artifacts_lineage/{name}][%d] RunsV1_GetRunArtifactLineage default  %+v", o._statusCode, o.Payload)
}

func (o *RunsV1GetRunArtifactLineageDefault) GetPayload() *service_model.RuntimeError {
	return o.Payload
}

func (o *RunsV1GetRunArtifactLineageDefault) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(service_model.RuntimeError)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}