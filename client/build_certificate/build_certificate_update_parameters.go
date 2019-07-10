// Code generated by go-swagger; DO NOT EDIT.

package build_certificate

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

// NewBuildCertificateUpdateParams creates a new BuildCertificateUpdateParams object
// with the default values initialized.
func NewBuildCertificateUpdateParams() *BuildCertificateUpdateParams {
	var ()
	return &BuildCertificateUpdateParams{

		timeout: cr.DefaultTimeout,
	}
}

// NewBuildCertificateUpdateParamsWithTimeout creates a new BuildCertificateUpdateParams object
// with the default values initialized, and the ability to set a timeout on a request
func NewBuildCertificateUpdateParamsWithTimeout(timeout time.Duration) *BuildCertificateUpdateParams {
	var ()
	return &BuildCertificateUpdateParams{

		timeout: timeout,
	}
}

// NewBuildCertificateUpdateParamsWithContext creates a new BuildCertificateUpdateParams object
// with the default values initialized, and the ability to set a context for a request
func NewBuildCertificateUpdateParamsWithContext(ctx context.Context) *BuildCertificateUpdateParams {
	var ()
	return &BuildCertificateUpdateParams{

		Context: ctx,
	}
}

// NewBuildCertificateUpdateParamsWithHTTPClient creates a new BuildCertificateUpdateParams object
// with the default values initialized, and the ability to set a custom HTTPClient for a request
func NewBuildCertificateUpdateParamsWithHTTPClient(client *http.Client) *BuildCertificateUpdateParams {
	var ()
	return &BuildCertificateUpdateParams{
		HTTPClient: client,
	}
}

/*BuildCertificateUpdateParams contains all the parameters to send to the API endpoint
for the build certificate update operation typically these are written to a http.Request
*/
type BuildCertificateUpdateParams struct {

	/*AppSlug
	  App slug

	*/
	AppSlug string
	/*BuildCertificate
	  Build certificate parameters

	*/
	BuildCertificate *models.V0BuildCertificateUpdateParams
	/*BuildCertificateSlug
	  Build certificate slug

	*/
	BuildCertificateSlug string

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithTimeout adds the timeout to the build certificate update params
func (o *BuildCertificateUpdateParams) WithTimeout(timeout time.Duration) *BuildCertificateUpdateParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the build certificate update params
func (o *BuildCertificateUpdateParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the build certificate update params
func (o *BuildCertificateUpdateParams) WithContext(ctx context.Context) *BuildCertificateUpdateParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the build certificate update params
func (o *BuildCertificateUpdateParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the build certificate update params
func (o *BuildCertificateUpdateParams) WithHTTPClient(client *http.Client) *BuildCertificateUpdateParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the build certificate update params
func (o *BuildCertificateUpdateParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithAppSlug adds the appSlug to the build certificate update params
func (o *BuildCertificateUpdateParams) WithAppSlug(appSlug string) *BuildCertificateUpdateParams {
	o.SetAppSlug(appSlug)
	return o
}

// SetAppSlug adds the appSlug to the build certificate update params
func (o *BuildCertificateUpdateParams) SetAppSlug(appSlug string) {
	o.AppSlug = appSlug
}

// WithBuildCertificate adds the buildCertificate to the build certificate update params
func (o *BuildCertificateUpdateParams) WithBuildCertificate(buildCertificate *models.V0BuildCertificateUpdateParams) *BuildCertificateUpdateParams {
	o.SetBuildCertificate(buildCertificate)
	return o
}

// SetBuildCertificate adds the buildCertificate to the build certificate update params
func (o *BuildCertificateUpdateParams) SetBuildCertificate(buildCertificate *models.V0BuildCertificateUpdateParams) {
	o.BuildCertificate = buildCertificate
}

// WithBuildCertificateSlug adds the buildCertificateSlug to the build certificate update params
func (o *BuildCertificateUpdateParams) WithBuildCertificateSlug(buildCertificateSlug string) *BuildCertificateUpdateParams {
	o.SetBuildCertificateSlug(buildCertificateSlug)
	return o
}

// SetBuildCertificateSlug adds the buildCertificateSlug to the build certificate update params
func (o *BuildCertificateUpdateParams) SetBuildCertificateSlug(buildCertificateSlug string) {
	o.BuildCertificateSlug = buildCertificateSlug
}

// WriteToRequest writes these params to a swagger request
func (o *BuildCertificateUpdateParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error

	// path param app-slug
	if err := r.SetPathParam("app-slug", o.AppSlug); err != nil {
		return err
	}

	if o.BuildCertificate != nil {
		if err := r.SetBodyParam(o.BuildCertificate); err != nil {
			return err
		}
	}

	// path param build-certificate-slug
	if err := r.SetPathParam("build-certificate-slug", o.BuildCertificateSlug); err != nil {
		return err
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}
