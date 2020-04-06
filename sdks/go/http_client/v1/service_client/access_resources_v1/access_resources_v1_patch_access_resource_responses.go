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

package access_resources_v1

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"
	"io"

	"github.com/go-openapi/runtime"
	"github.com/go-openapi/strfmt"

	"github.com/polyaxon/polyaxon/sdks/go/http_client/v1/service_model"
)

// AccessResourcesV1PatchAccessResourceReader is a Reader for the AccessResourcesV1PatchAccessResource structure.
type AccessResourcesV1PatchAccessResourceReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *AccessResourcesV1PatchAccessResourceReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewAccessResourcesV1PatchAccessResourceOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 204:
		result := NewAccessResourcesV1PatchAccessResourceNoContent()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 403:
		result := NewAccessResourcesV1PatchAccessResourceForbidden()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 404:
		result := NewAccessResourcesV1PatchAccessResourceNotFound()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	default:
		result := NewAccessResourcesV1PatchAccessResourceDefault(response.Code())
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		if response.Code()/100 == 2 {
			return result, nil
		}
		return nil, result
	}
}

// NewAccessResourcesV1PatchAccessResourceOK creates a AccessResourcesV1PatchAccessResourceOK with default headers values
func NewAccessResourcesV1PatchAccessResourceOK() *AccessResourcesV1PatchAccessResourceOK {
	return &AccessResourcesV1PatchAccessResourceOK{}
}

/*AccessResourcesV1PatchAccessResourceOK handles this case with default header values.

A successful response.
*/
type AccessResourcesV1PatchAccessResourceOK struct {
	Payload *service_model.V1AccessResource
}

func (o *AccessResourcesV1PatchAccessResourceOK) Error() string {
	return fmt.Sprintf("[PATCH /api/v1/orgs/{owner}/access_resources/{access_resource.uuid}][%d] accessResourcesV1PatchAccessResourceOK  %+v", 200, o.Payload)
}

func (o *AccessResourcesV1PatchAccessResourceOK) GetPayload() *service_model.V1AccessResource {
	return o.Payload
}

func (o *AccessResourcesV1PatchAccessResourceOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(service_model.V1AccessResource)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewAccessResourcesV1PatchAccessResourceNoContent creates a AccessResourcesV1PatchAccessResourceNoContent with default headers values
func NewAccessResourcesV1PatchAccessResourceNoContent() *AccessResourcesV1PatchAccessResourceNoContent {
	return &AccessResourcesV1PatchAccessResourceNoContent{}
}

/*AccessResourcesV1PatchAccessResourceNoContent handles this case with default header values.

No content.
*/
type AccessResourcesV1PatchAccessResourceNoContent struct {
	Payload interface{}
}

func (o *AccessResourcesV1PatchAccessResourceNoContent) Error() string {
	return fmt.Sprintf("[PATCH /api/v1/orgs/{owner}/access_resources/{access_resource.uuid}][%d] accessResourcesV1PatchAccessResourceNoContent  %+v", 204, o.Payload)
}

func (o *AccessResourcesV1PatchAccessResourceNoContent) GetPayload() interface{} {
	return o.Payload
}

func (o *AccessResourcesV1PatchAccessResourceNoContent) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// response payload
	if err := consumer.Consume(response.Body(), &o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewAccessResourcesV1PatchAccessResourceForbidden creates a AccessResourcesV1PatchAccessResourceForbidden with default headers values
func NewAccessResourcesV1PatchAccessResourceForbidden() *AccessResourcesV1PatchAccessResourceForbidden {
	return &AccessResourcesV1PatchAccessResourceForbidden{}
}

/*AccessResourcesV1PatchAccessResourceForbidden handles this case with default header values.

You don't have permission to access the resource.
*/
type AccessResourcesV1PatchAccessResourceForbidden struct {
	Payload interface{}
}

func (o *AccessResourcesV1PatchAccessResourceForbidden) Error() string {
	return fmt.Sprintf("[PATCH /api/v1/orgs/{owner}/access_resources/{access_resource.uuid}][%d] accessResourcesV1PatchAccessResourceForbidden  %+v", 403, o.Payload)
}

func (o *AccessResourcesV1PatchAccessResourceForbidden) GetPayload() interface{} {
	return o.Payload
}

func (o *AccessResourcesV1PatchAccessResourceForbidden) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// response payload
	if err := consumer.Consume(response.Body(), &o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewAccessResourcesV1PatchAccessResourceNotFound creates a AccessResourcesV1PatchAccessResourceNotFound with default headers values
func NewAccessResourcesV1PatchAccessResourceNotFound() *AccessResourcesV1PatchAccessResourceNotFound {
	return &AccessResourcesV1PatchAccessResourceNotFound{}
}

/*AccessResourcesV1PatchAccessResourceNotFound handles this case with default header values.

Resource does not exist.
*/
type AccessResourcesV1PatchAccessResourceNotFound struct {
	Payload interface{}
}

func (o *AccessResourcesV1PatchAccessResourceNotFound) Error() string {
	return fmt.Sprintf("[PATCH /api/v1/orgs/{owner}/access_resources/{access_resource.uuid}][%d] accessResourcesV1PatchAccessResourceNotFound  %+v", 404, o.Payload)
}

func (o *AccessResourcesV1PatchAccessResourceNotFound) GetPayload() interface{} {
	return o.Payload
}

func (o *AccessResourcesV1PatchAccessResourceNotFound) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// response payload
	if err := consumer.Consume(response.Body(), &o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewAccessResourcesV1PatchAccessResourceDefault creates a AccessResourcesV1PatchAccessResourceDefault with default headers values
func NewAccessResourcesV1PatchAccessResourceDefault(code int) *AccessResourcesV1PatchAccessResourceDefault {
	return &AccessResourcesV1PatchAccessResourceDefault{
		_statusCode: code,
	}
}

/*AccessResourcesV1PatchAccessResourceDefault handles this case with default header values.

An unexpected error response
*/
type AccessResourcesV1PatchAccessResourceDefault struct {
	_statusCode int

	Payload *service_model.RuntimeError
}

// Code gets the status code for the access resources v1 patch access resource default response
func (o *AccessResourcesV1PatchAccessResourceDefault) Code() int {
	return o._statusCode
}

func (o *AccessResourcesV1PatchAccessResourceDefault) Error() string {
	return fmt.Sprintf("[PATCH /api/v1/orgs/{owner}/access_resources/{access_resource.uuid}][%d] AccessResourcesV1_PatchAccessResource default  %+v", o._statusCode, o.Payload)
}

func (o *AccessResourcesV1PatchAccessResourceDefault) GetPayload() *service_model.RuntimeError {
	return o.Payload
}

func (o *AccessResourcesV1PatchAccessResourceDefault) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(service_model.RuntimeError)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}