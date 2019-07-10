// Code generated by go-swagger; DO NOT EDIT.

package app_setup

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

// NewAppFinishParams creates a new AppFinishParams object
// with the default values initialized.
func NewAppFinishParams() *AppFinishParams {
	var ()
	return &AppFinishParams{

		timeout: cr.DefaultTimeout,
	}
}

// NewAppFinishParamsWithTimeout creates a new AppFinishParams object
// with the default values initialized, and the ability to set a timeout on a request
func NewAppFinishParamsWithTimeout(timeout time.Duration) *AppFinishParams {
	var ()
	return &AppFinishParams{

		timeout: timeout,
	}
}

// NewAppFinishParamsWithContext creates a new AppFinishParams object
// with the default values initialized, and the ability to set a context for a request
func NewAppFinishParamsWithContext(ctx context.Context) *AppFinishParams {
	var ()
	return &AppFinishParams{

		Context: ctx,
	}
}

// NewAppFinishParamsWithHTTPClient creates a new AppFinishParams object
// with the default values initialized, and the ability to set a custom HTTPClient for a request
func NewAppFinishParamsWithHTTPClient(client *http.Client) *AppFinishParams {
	var ()
	return &AppFinishParams{
		HTTPClient: client,
	}
}

/*AppFinishParams contains all the parameters to send to the API endpoint
for the app finish operation typically these are written to a http.Request
*/
type AppFinishParams struct {

	/*App
	  App finish parameters

	*/
	App *models.V0AppFinishParams
	/*AppSlug
	  App slug

	*/
	AppSlug string

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithTimeout adds the timeout to the app finish params
func (o *AppFinishParams) WithTimeout(timeout time.Duration) *AppFinishParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the app finish params
func (o *AppFinishParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the app finish params
func (o *AppFinishParams) WithContext(ctx context.Context) *AppFinishParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the app finish params
func (o *AppFinishParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the app finish params
func (o *AppFinishParams) WithHTTPClient(client *http.Client) *AppFinishParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the app finish params
func (o *AppFinishParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithApp adds the app to the app finish params
func (o *AppFinishParams) WithApp(app *models.V0AppFinishParams) *AppFinishParams {
	o.SetApp(app)
	return o
}

// SetApp adds the app to the app finish params
func (o *AppFinishParams) SetApp(app *models.V0AppFinishParams) {
	o.App = app
}

// WithAppSlug adds the appSlug to the app finish params
func (o *AppFinishParams) WithAppSlug(appSlug string) *AppFinishParams {
	o.SetAppSlug(appSlug)
	return o
}

// SetAppSlug adds the appSlug to the app finish params
func (o *AppFinishParams) SetAppSlug(appSlug string) {
	o.AppSlug = appSlug
}

// WriteToRequest writes these params to a swagger request
func (o *AppFinishParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error

	if o.App != nil {
		if err := r.SetBodyParam(o.App); err != nil {
			return err
		}
	}

	// path param app-slug
	if err := r.SetPathParam("app-slug", o.AppSlug); err != nil {
		return err
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}
