// Code generated by go-swagger; DO NOT EDIT.

package build_request

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

	models "github.com/peartherapeutics/bitrise-step-api-build-trigger/models"
)

// NewBuildRequestUpdateParams creates a new BuildRequestUpdateParams object
// with the default values initialized.
func NewBuildRequestUpdateParams() *BuildRequestUpdateParams {
	var ()
	return &BuildRequestUpdateParams{

		timeout: cr.DefaultTimeout,
	}
}

// NewBuildRequestUpdateParamsWithTimeout creates a new BuildRequestUpdateParams object
// with the default values initialized, and the ability to set a timeout on a request
func NewBuildRequestUpdateParamsWithTimeout(timeout time.Duration) *BuildRequestUpdateParams {
	var ()
	return &BuildRequestUpdateParams{

		timeout: timeout,
	}
}

// NewBuildRequestUpdateParamsWithContext creates a new BuildRequestUpdateParams object
// with the default values initialized, and the ability to set a context for a request
func NewBuildRequestUpdateParamsWithContext(ctx context.Context) *BuildRequestUpdateParams {
	var ()
	return &BuildRequestUpdateParams{

		Context: ctx,
	}
}

// NewBuildRequestUpdateParamsWithHTTPClient creates a new BuildRequestUpdateParams object
// with the default values initialized, and the ability to set a custom HTTPClient for a request
func NewBuildRequestUpdateParamsWithHTTPClient(client *http.Client) *BuildRequestUpdateParams {
	var ()
	return &BuildRequestUpdateParams{
		HTTPClient: client,
	}
}

/*BuildRequestUpdateParams contains all the parameters to send to the API endpoint
for the build request update operation typically these are written to a http.Request
*/
type BuildRequestUpdateParams struct {

	/*AppSlug
	  App slug

	*/
	AppSlug string
	/*BuildRequest
	  Build request parameters

	*/
	BuildRequest *models.V0BuildRequestUpdateParams
	/*BuildRequestSlug
	  Build request slug

	*/
	BuildRequestSlug string

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithTimeout adds the timeout to the build request update params
func (o *BuildRequestUpdateParams) WithTimeout(timeout time.Duration) *BuildRequestUpdateParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the build request update params
func (o *BuildRequestUpdateParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the build request update params
func (o *BuildRequestUpdateParams) WithContext(ctx context.Context) *BuildRequestUpdateParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the build request update params
func (o *BuildRequestUpdateParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the build request update params
func (o *BuildRequestUpdateParams) WithHTTPClient(client *http.Client) *BuildRequestUpdateParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the build request update params
func (o *BuildRequestUpdateParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithAppSlug adds the appSlug to the build request update params
func (o *BuildRequestUpdateParams) WithAppSlug(appSlug string) *BuildRequestUpdateParams {
	o.SetAppSlug(appSlug)
	return o
}

// SetAppSlug adds the appSlug to the build request update params
func (o *BuildRequestUpdateParams) SetAppSlug(appSlug string) {
	o.AppSlug = appSlug
}

// WithBuildRequest adds the buildRequest to the build request update params
func (o *BuildRequestUpdateParams) WithBuildRequest(buildRequest *models.V0BuildRequestUpdateParams) *BuildRequestUpdateParams {
	o.SetBuildRequest(buildRequest)
	return o
}

// SetBuildRequest adds the buildRequest to the build request update params
func (o *BuildRequestUpdateParams) SetBuildRequest(buildRequest *models.V0BuildRequestUpdateParams) {
	o.BuildRequest = buildRequest
}

// WithBuildRequestSlug adds the buildRequestSlug to the build request update params
func (o *BuildRequestUpdateParams) WithBuildRequestSlug(buildRequestSlug string) *BuildRequestUpdateParams {
	o.SetBuildRequestSlug(buildRequestSlug)
	return o
}

// SetBuildRequestSlug adds the buildRequestSlug to the build request update params
func (o *BuildRequestUpdateParams) SetBuildRequestSlug(buildRequestSlug string) {
	o.BuildRequestSlug = buildRequestSlug
}

// WriteToRequest writes these params to a swagger request
func (o *BuildRequestUpdateParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error

	// path param app-slug
	if err := r.SetPathParam("app-slug", o.AppSlug); err != nil {
		return err
	}

	if o.BuildRequest != nil {
		if err := r.SetBodyParam(o.BuildRequest); err != nil {
			return err
		}
	}

	// path param build-request-slug
	if err := r.SetPathParam("build-request-slug", o.BuildRequestSlug); err != nil {
		return err
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}
