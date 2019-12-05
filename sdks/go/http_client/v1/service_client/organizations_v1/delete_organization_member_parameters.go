// Copyright 2019 Polyaxon, Inc.
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

package organizations_v1

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"context"
	"net/http"
	"time"

	"github.com/go-openapi/errors"
	"github.com/go-openapi/runtime"
	cr "github.com/go-openapi/runtime/client"

	strfmt "github.com/go-openapi/strfmt"
)

// NewDeleteOrganizationMemberParams creates a new DeleteOrganizationMemberParams object
// with the default values initialized.
func NewDeleteOrganizationMemberParams() *DeleteOrganizationMemberParams {
	var ()
	return &DeleteOrganizationMemberParams{

		timeout: cr.DefaultTimeout,
	}
}

// NewDeleteOrganizationMemberParamsWithTimeout creates a new DeleteOrganizationMemberParams object
// with the default values initialized, and the ability to set a timeout on a request
func NewDeleteOrganizationMemberParamsWithTimeout(timeout time.Duration) *DeleteOrganizationMemberParams {
	var ()
	return &DeleteOrganizationMemberParams{

		timeout: timeout,
	}
}

// NewDeleteOrganizationMemberParamsWithContext creates a new DeleteOrganizationMemberParams object
// with the default values initialized, and the ability to set a context for a request
func NewDeleteOrganizationMemberParamsWithContext(ctx context.Context) *DeleteOrganizationMemberParams {
	var ()
	return &DeleteOrganizationMemberParams{

		Context: ctx,
	}
}

// NewDeleteOrganizationMemberParamsWithHTTPClient creates a new DeleteOrganizationMemberParams object
// with the default values initialized, and the ability to set a custom HTTPClient for a request
func NewDeleteOrganizationMemberParamsWithHTTPClient(client *http.Client) *DeleteOrganizationMemberParams {
	var ()
	return &DeleteOrganizationMemberParams{
		HTTPClient: client,
	}
}

/*DeleteOrganizationMemberParams contains all the parameters to send to the API endpoint
for the delete organization member operation typically these are written to a http.Request
*/
type DeleteOrganizationMemberParams struct {

	/*MemberCreatedAt
	  Optional time when the entityt was created.

	*/
	MemberCreatedAt *strfmt.DateTime
	/*MemberRole
	  Role.

	*/
	MemberRole *string
	/*MemberUpdatedAt
	  Optional last time the entity was updated.

	*/
	MemberUpdatedAt *strfmt.DateTime
	/*MemberUser
	  User

	*/
	MemberUser string
	/*Owner
	  Owner of the namespace

	*/
	Owner string

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithTimeout adds the timeout to the delete organization member params
func (o *DeleteOrganizationMemberParams) WithTimeout(timeout time.Duration) *DeleteOrganizationMemberParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the delete organization member params
func (o *DeleteOrganizationMemberParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the delete organization member params
func (o *DeleteOrganizationMemberParams) WithContext(ctx context.Context) *DeleteOrganizationMemberParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the delete organization member params
func (o *DeleteOrganizationMemberParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the delete organization member params
func (o *DeleteOrganizationMemberParams) WithHTTPClient(client *http.Client) *DeleteOrganizationMemberParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the delete organization member params
func (o *DeleteOrganizationMemberParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithMemberCreatedAt adds the memberCreatedAt to the delete organization member params
func (o *DeleteOrganizationMemberParams) WithMemberCreatedAt(memberCreatedAt *strfmt.DateTime) *DeleteOrganizationMemberParams {
	o.SetMemberCreatedAt(memberCreatedAt)
	return o
}

// SetMemberCreatedAt adds the memberCreatedAt to the delete organization member params
func (o *DeleteOrganizationMemberParams) SetMemberCreatedAt(memberCreatedAt *strfmt.DateTime) {
	o.MemberCreatedAt = memberCreatedAt
}

// WithMemberRole adds the memberRole to the delete organization member params
func (o *DeleteOrganizationMemberParams) WithMemberRole(memberRole *string) *DeleteOrganizationMemberParams {
	o.SetMemberRole(memberRole)
	return o
}

// SetMemberRole adds the memberRole to the delete organization member params
func (o *DeleteOrganizationMemberParams) SetMemberRole(memberRole *string) {
	o.MemberRole = memberRole
}

// WithMemberUpdatedAt adds the memberUpdatedAt to the delete organization member params
func (o *DeleteOrganizationMemberParams) WithMemberUpdatedAt(memberUpdatedAt *strfmt.DateTime) *DeleteOrganizationMemberParams {
	o.SetMemberUpdatedAt(memberUpdatedAt)
	return o
}

// SetMemberUpdatedAt adds the memberUpdatedAt to the delete organization member params
func (o *DeleteOrganizationMemberParams) SetMemberUpdatedAt(memberUpdatedAt *strfmt.DateTime) {
	o.MemberUpdatedAt = memberUpdatedAt
}

// WithMemberUser adds the memberUser to the delete organization member params
func (o *DeleteOrganizationMemberParams) WithMemberUser(memberUser string) *DeleteOrganizationMemberParams {
	o.SetMemberUser(memberUser)
	return o
}

// SetMemberUser adds the memberUser to the delete organization member params
func (o *DeleteOrganizationMemberParams) SetMemberUser(memberUser string) {
	o.MemberUser = memberUser
}

// WithOwner adds the owner to the delete organization member params
func (o *DeleteOrganizationMemberParams) WithOwner(owner string) *DeleteOrganizationMemberParams {
	o.SetOwner(owner)
	return o
}

// SetOwner adds the owner to the delete organization member params
func (o *DeleteOrganizationMemberParams) SetOwner(owner string) {
	o.Owner = owner
}

// WriteToRequest writes these params to a swagger request
func (o *DeleteOrganizationMemberParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error

	if o.MemberCreatedAt != nil {

		// query param member.created_at
		var qrMemberCreatedAt strfmt.DateTime
		if o.MemberCreatedAt != nil {
			qrMemberCreatedAt = *o.MemberCreatedAt
		}
		qMemberCreatedAt := qrMemberCreatedAt.String()
		if qMemberCreatedAt != "" {
			if err := r.SetQueryParam("member.created_at", qMemberCreatedAt); err != nil {
				return err
			}
		}

	}

	if o.MemberRole != nil {

		// query param member.role
		var qrMemberRole string
		if o.MemberRole != nil {
			qrMemberRole = *o.MemberRole
		}
		qMemberRole := qrMemberRole
		if qMemberRole != "" {
			if err := r.SetQueryParam("member.role", qMemberRole); err != nil {
				return err
			}
		}

	}

	if o.MemberUpdatedAt != nil {

		// query param member.updated_at
		var qrMemberUpdatedAt strfmt.DateTime
		if o.MemberUpdatedAt != nil {
			qrMemberUpdatedAt = *o.MemberUpdatedAt
		}
		qMemberUpdatedAt := qrMemberUpdatedAt.String()
		if qMemberUpdatedAt != "" {
			if err := r.SetQueryParam("member.updated_at", qMemberUpdatedAt); err != nil {
				return err
			}
		}

	}

	// path param member.user
	if err := r.SetPathParam("member.user", o.MemberUser); err != nil {
		return err
	}

	// path param owner
	if err := r.SetPathParam("owner", o.Owner); err != nil {
		return err
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}