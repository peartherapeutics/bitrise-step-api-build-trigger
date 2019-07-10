// Code generated by go-swagger; DO NOT EDIT.

package provisioning_profile

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

// NewProvisioningProfileDeleteParams creates a new ProvisioningProfileDeleteParams object
// with the default values initialized.
func NewProvisioningProfileDeleteParams() *ProvisioningProfileDeleteParams {
	var ()
	return &ProvisioningProfileDeleteParams{

		timeout: cr.DefaultTimeout,
	}
}

// NewProvisioningProfileDeleteParamsWithTimeout creates a new ProvisioningProfileDeleteParams object
// with the default values initialized, and the ability to set a timeout on a request
func NewProvisioningProfileDeleteParamsWithTimeout(timeout time.Duration) *ProvisioningProfileDeleteParams {
	var ()
	return &ProvisioningProfileDeleteParams{

		timeout: timeout,
	}
}

// NewProvisioningProfileDeleteParamsWithContext creates a new ProvisioningProfileDeleteParams object
// with the default values initialized, and the ability to set a context for a request
func NewProvisioningProfileDeleteParamsWithContext(ctx context.Context) *ProvisioningProfileDeleteParams {
	var ()
	return &ProvisioningProfileDeleteParams{

		Context: ctx,
	}
}

// NewProvisioningProfileDeleteParamsWithHTTPClient creates a new ProvisioningProfileDeleteParams object
// with the default values initialized, and the ability to set a custom HTTPClient for a request
func NewProvisioningProfileDeleteParamsWithHTTPClient(client *http.Client) *ProvisioningProfileDeleteParams {
	var ()
	return &ProvisioningProfileDeleteParams{
		HTTPClient: client,
	}
}

/*ProvisioningProfileDeleteParams contains all the parameters to send to the API endpoint
for the provisioning profile delete operation typically these are written to a http.Request
*/
type ProvisioningProfileDeleteParams struct {

	/*AppSlug
	  App slug

	*/
	AppSlug string
	/*ProvisioningProfileSlug
	  Provisioning profile slug

	*/
	ProvisioningProfileSlug string

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithTimeout adds the timeout to the provisioning profile delete params
func (o *ProvisioningProfileDeleteParams) WithTimeout(timeout time.Duration) *ProvisioningProfileDeleteParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the provisioning profile delete params
func (o *ProvisioningProfileDeleteParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the provisioning profile delete params
func (o *ProvisioningProfileDeleteParams) WithContext(ctx context.Context) *ProvisioningProfileDeleteParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the provisioning profile delete params
func (o *ProvisioningProfileDeleteParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the provisioning profile delete params
func (o *ProvisioningProfileDeleteParams) WithHTTPClient(client *http.Client) *ProvisioningProfileDeleteParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the provisioning profile delete params
func (o *ProvisioningProfileDeleteParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithAppSlug adds the appSlug to the provisioning profile delete params
func (o *ProvisioningProfileDeleteParams) WithAppSlug(appSlug string) *ProvisioningProfileDeleteParams {
	o.SetAppSlug(appSlug)
	return o
}

// SetAppSlug adds the appSlug to the provisioning profile delete params
func (o *ProvisioningProfileDeleteParams) SetAppSlug(appSlug string) {
	o.AppSlug = appSlug
}

// WithProvisioningProfileSlug adds the provisioningProfileSlug to the provisioning profile delete params
func (o *ProvisioningProfileDeleteParams) WithProvisioningProfileSlug(provisioningProfileSlug string) *ProvisioningProfileDeleteParams {
	o.SetProvisioningProfileSlug(provisioningProfileSlug)
	return o
}

// SetProvisioningProfileSlug adds the provisioningProfileSlug to the provisioning profile delete params
func (o *ProvisioningProfileDeleteParams) SetProvisioningProfileSlug(provisioningProfileSlug string) {
	o.ProvisioningProfileSlug = provisioningProfileSlug
}

// WriteToRequest writes these params to a swagger request
func (o *ProvisioningProfileDeleteParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error

	// path param app-slug
	if err := r.SetPathParam("app-slug", o.AppSlug); err != nil {
		return err
	}

	// path param provisioning-profile-slug
	if err := r.SetPathParam("provisioning-profile-slug", o.ProvisioningProfileSlug); err != nil {
		return err
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}
