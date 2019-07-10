// Code generated by go-swagger; DO NOT EDIT.

package builds

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"github.com/go-openapi/runtime"

	strfmt "github.com/go-openapi/strfmt"
)

// New creates a new builds API client.
func New(transport runtime.ClientTransport, formats strfmt.Registry) *Client {
	return &Client{transport: transport, formats: formats}
}

/*
Client for builds API
*/
type Client struct {
	transport runtime.ClientTransport
	formats   strfmt.Registry
}

/*
BuildAbort aborts a specific build

Abort a specific build. Set an abort reason with the `abort_reason` parameter. Use the `abort_with_success` parameter to abort a build but still count it as a successful one.
*/
func (a *Client) BuildAbort(params *BuildAbortParams, authInfo runtime.ClientAuthInfoWriter) (*BuildAbortOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewBuildAbortParams()
	}

	result, err := a.transport.Submit(&runtime.ClientOperation{
		ID:                 "build-abort",
		Method:             "POST",
		PathPattern:        "/apps/{app-slug}/builds/{build-slug}/abort",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"https"},
		Params:             params,
		Reader:             &BuildAbortReader{formats: a.formats},
		AuthInfo:           authInfo,
		Context:            params.Context,
		Client:             params.HTTPClient,
	})
	if err != nil {
		return nil, err
	}
	return result.(*BuildAbortOK), nil

}

/*
BuildBitriseYmlShow gets the bitrise yml of a build

Get the bitrise.yml file of one of the builds of a given app. This will return the `bitrise.yml` configuration with which the build ran. You can compare it to [the current bitrise.yml configuration](https://api-docs.bitrise.io/#/application/app-config-datastore-show) of the app.
*/
func (a *Client) BuildBitriseYmlShow(params *BuildBitriseYmlShowParams, authInfo runtime.ClientAuthInfoWriter) (*BuildBitriseYmlShowOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewBuildBitriseYmlShowParams()
	}

	result, err := a.transport.Submit(&runtime.ClientOperation{
		ID:                 "build-bitrise-yml-show",
		Method:             "GET",
		PathPattern:        "/apps/{app-slug}/builds/{build-slug}/bitrise.yml",
		ProducesMediaTypes: []string{"text/plain"},
		ConsumesMediaTypes: []string{""},
		Schemes:            []string{"https"},
		Params:             params,
		Reader:             &BuildBitriseYmlShowReader{formats: a.formats},
		AuthInfo:           authInfo,
		Context:            params.Context,
		Client:             params.HTTPClient,
	})
	if err != nil {
		return nil, err
	}
	return result.(*BuildBitriseYmlShowOK), nil

}

/*
BuildList lists all builds of an app

List all the builds of a specified Bitrise app. Set parameters to filter builds: for example, you can search for builds run with a given workflow or all builds that were triggered by Pull Requests. It returns all the relevant data of the build.
*/
func (a *Client) BuildList(params *BuildListParams, authInfo runtime.ClientAuthInfoWriter) (*BuildListOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewBuildListParams()
	}

	result, err := a.transport.Submit(&runtime.ClientOperation{
		ID:                 "build-list",
		Method:             "GET",
		PathPattern:        "/apps/{app-slug}/builds",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{""},
		Schemes:            []string{"https"},
		Params:             params,
		Reader:             &BuildListReader{formats: a.formats},
		AuthInfo:           authInfo,
		Context:            params.Context,
		Client:             params.HTTPClient,
	})
	if err != nil {
		return nil, err
	}
	return result.(*BuildListOK), nil

}

/*
BuildListAll lists all builds

List all the Bitrise builds that can be accessed with the authenticated account. Filter builds based on their owner, using the owner slug, or the status of the build.
*/
func (a *Client) BuildListAll(params *BuildListAllParams, authInfo runtime.ClientAuthInfoWriter) (*BuildListAllOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewBuildListAllParams()
	}

	result, err := a.transport.Submit(&runtime.ClientOperation{
		ID:                 "build-list-all",
		Method:             "GET",
		PathPattern:        "/builds",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{""},
		Schemes:            []string{"https"},
		Params:             params,
		Reader:             &BuildListAllReader{formats: a.formats},
		AuthInfo:           authInfo,
		Context:            params.Context,
		Client:             params.HTTPClient,
	})
	if err != nil {
		return nil, err
	}
	return result.(*BuildListAllOK), nil

}

/*
BuildLog gets the build log of a build

Get the build log of a specified build of a Bitrise app. You can get the build slug either by calling the [/builds](https://api-docs.bitrise.io/#/builds/build-list) endpoint or by clicking on the build on bitrise.io and copying the slug from the URL.
*/
func (a *Client) BuildLog(params *BuildLogParams, authInfo runtime.ClientAuthInfoWriter) (*BuildLogOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewBuildLogParams()
	}

	result, err := a.transport.Submit(&runtime.ClientOperation{
		ID:                 "build-log",
		Method:             "GET",
		PathPattern:        "/apps/{app-slug}/builds/{build-slug}/log",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{""},
		Schemes:            []string{"https"},
		Params:             params,
		Reader:             &BuildLogReader{formats: a.formats},
		AuthInfo:           authInfo,
		Context:            params.Context,
		Client:             params.HTTPClient,
	})
	if err != nil {
		return nil, err
	}
	return result.(*BuildLogOK), nil

}

/*
BuildShow gets a build of a given app

Get the specified build of a given Bitrise app. You need to provide both an app slug and a build slug. You can get the build slug either by calling the [/builds](https://api-docs.bitrise.io/#/builds/build-list) endpoint or by clicking on the build on bitrise.io and copying the slug from the URL. The endpoint returns all the relevant data of the build.
*/
func (a *Client) BuildShow(params *BuildShowParams, authInfo runtime.ClientAuthInfoWriter) (*BuildShowOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewBuildShowParams()
	}

	result, err := a.transport.Submit(&runtime.ClientOperation{
		ID:                 "build-show",
		Method:             "GET",
		PathPattern:        "/apps/{app-slug}/builds/{build-slug}",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{""},
		Schemes:            []string{"https"},
		Params:             params,
		Reader:             &BuildShowReader{formats: a.formats},
		AuthInfo:           authInfo,
		Context:            params.Context,
		Client:             params.HTTPClient,
	})
	if err != nil {
		return nil, err
	}
	return result.(*BuildShowOK), nil

}

/*
BuildTrigger triggers a new build

Trigger a new build. Specify an app slug and at least one parameter out of three: a git tag or git commit hash, a branch, or a workflow ID. You can also set specific parameters for Pull Request builds and define additional environment variables for your build. [Check out our detailed guide](https://devcenter.bitrise.io/api/build-trigger/).
*/
func (a *Client) BuildTrigger(params *BuildTriggerParams, authInfo runtime.ClientAuthInfoWriter) (*BuildTriggerCreated, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewBuildTriggerParams()
	}

	result, err := a.transport.Submit(&runtime.ClientOperation{
		ID:                 "build-trigger",
		Method:             "POST",
		PathPattern:        "/apps/{app-slug}/builds",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"https"},
		Params:             params,
		Reader:             &BuildTriggerReader{formats: a.formats},
		AuthInfo:           authInfo,
		Context:            params.Context,
		Client:             params.HTTPClient,
	})
	if err != nil {
		return nil, err
	}
	return result.(*BuildTriggerCreated), nil

}

/*
BuildWorkflowList lists the workflows of an app

List the workflows that were triggered at any time for a given Bitrise app. Note that it might list workflows that are currently not defined in the app's `bitrise.yml` configuration - and conversely, workflows that were never triggered will not be listed even if they are defined in the `bitrise.yml` file.
*/
func (a *Client) BuildWorkflowList(params *BuildWorkflowListParams, authInfo runtime.ClientAuthInfoWriter) (*BuildWorkflowListOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewBuildWorkflowListParams()
	}

	result, err := a.transport.Submit(&runtime.ClientOperation{
		ID:                 "build-workflow-list",
		Method:             "GET",
		PathPattern:        "/apps/{app-slug}/build-workflows",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{""},
		Schemes:            []string{"https"},
		Params:             params,
		Reader:             &BuildWorkflowListReader{formats: a.formats},
		AuthInfo:           authInfo,
		Context:            params.Context,
		Client:             params.HTTPClient,
	})
	if err != nil {
		return nil, err
	}
	return result.(*BuildWorkflowListOK), nil

}

// SetTransport changes the transport on the client
func (a *Client) SetTransport(transport runtime.ClientTransport) {
	a.transport = transport
}
