// Code generated by go-swagger; DO NOT EDIT.

package access_control

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

// NewGetRolePermissionsParams creates a new GetRolePermissionsParams object
// with the default values initialized.
func NewGetRolePermissionsParams() *GetRolePermissionsParams {
	var ()
	return &GetRolePermissionsParams{

		timeout: cr.DefaultTimeout,
	}
}

// NewGetRolePermissionsParamsWithTimeout creates a new GetRolePermissionsParams object
// with the default values initialized, and the ability to set a timeout on a request
func NewGetRolePermissionsParamsWithTimeout(timeout time.Duration) *GetRolePermissionsParams {
	var ()
	return &GetRolePermissionsParams{

		timeout: timeout,
	}
}

// NewGetRolePermissionsParamsWithContext creates a new GetRolePermissionsParams object
// with the default values initialized, and the ability to set a context for a request
func NewGetRolePermissionsParamsWithContext(ctx context.Context) *GetRolePermissionsParams {
	var ()
	return &GetRolePermissionsParams{

		Context: ctx,
	}
}

// NewGetRolePermissionsParamsWithHTTPClient creates a new GetRolePermissionsParams object
// with the default values initialized, and the ability to set a custom HTTPClient for a request
func NewGetRolePermissionsParamsWithHTTPClient(client *http.Client) *GetRolePermissionsParams {
	var ()
	return &GetRolePermissionsParams{
		HTTPClient: client,
	}
}

/*GetRolePermissionsParams contains all the parameters to send to the API endpoint
for the get role permissions operation typically these are written to a http.Request
*/
type GetRolePermissionsParams struct {

	/*RoleID
	  The role ID to reference

	*/
	RoleID string

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithTimeout adds the timeout to the get role permissions params
func (o *GetRolePermissionsParams) WithTimeout(timeout time.Duration) *GetRolePermissionsParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the get role permissions params
func (o *GetRolePermissionsParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the get role permissions params
func (o *GetRolePermissionsParams) WithContext(ctx context.Context) *GetRolePermissionsParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the get role permissions params
func (o *GetRolePermissionsParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the get role permissions params
func (o *GetRolePermissionsParams) WithHTTPClient(client *http.Client) *GetRolePermissionsParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the get role permissions params
func (o *GetRolePermissionsParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithRoleID adds the roleID to the get role permissions params
func (o *GetRolePermissionsParams) WithRoleID(roleID string) *GetRolePermissionsParams {
	o.SetRoleID(roleID)
	return o
}

// SetRoleID adds the roleId to the get role permissions params
func (o *GetRolePermissionsParams) SetRoleID(roleID string) {
	o.RoleID = roleID
}

// WriteToRequest writes these params to a swagger request
func (o *GetRolePermissionsParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error

	// path param roleId
	if err := r.SetPathParam("roleId", o.RoleID); err != nil {
		return err
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}