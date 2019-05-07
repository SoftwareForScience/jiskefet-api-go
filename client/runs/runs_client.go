// Code generated by go-swagger; DO NOT EDIT.

package runs

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"github.com/go-openapi/runtime"

	strfmt "github.com/go-openapi/strfmt"
)

// New creates a new runs API client.
func New(transport runtime.ClientTransport, formats strfmt.Registry) *Client {
	return &Client{transport: transport, formats: formats}
}

/*
Client for runs API
*/
type Client struct {
	transport runtime.ClientTransport
	formats   strfmt.Registry
}

/*
GetRuns returns all runs
*/
func (a *Client) GetRuns(params *GetRunsParams, authInfo runtime.ClientAuthInfoWriter) (*GetRunsOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewGetRunsParams()
	}

	result, err := a.transport.Submit(&runtime.ClientOperation{
		ID:                 "GetRuns",
		Method:             "GET",
		PathPattern:        "/runs",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"http"},
		Params:             params,
		Reader:             &GetRunsReader{formats: a.formats},
		AuthInfo:           authInfo,
		Context:            params.Context,
		Client:             params.HTTPClient,
	})
	if err != nil {
		return nil, err
	}
	return result.(*GetRunsOK), nil

}

/*
GetRunsID returns a specific run
*/
func (a *Client) GetRunsID(params *GetRunsIDParams, authInfo runtime.ClientAuthInfoWriter) (*GetRunsIDOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewGetRunsIDParams()
	}

	result, err := a.transport.Submit(&runtime.ClientOperation{
		ID:                 "GetRunsID",
		Method:             "GET",
		PathPattern:        "/runs/{id}",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"http"},
		Params:             params,
		Reader:             &GetRunsIDReader{formats: a.formats},
		AuthInfo:           authInfo,
		Context:            params.Context,
		Client:             params.HTTPClient,
	})
	if err != nil {
		return nil, err
	}
	return result.(*GetRunsIDOK), nil

}

/*
PatchRunsID updates certain fields of a run
*/
func (a *Client) PatchRunsID(params *PatchRunsIDParams, authInfo runtime.ClientAuthInfoWriter) (*PatchRunsIDOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewPatchRunsIDParams()
	}

	result, err := a.transport.Submit(&runtime.ClientOperation{
		ID:                 "PatchRunsID",
		Method:             "PATCH",
		PathPattern:        "/runs/{id}",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"http"},
		Params:             params,
		Reader:             &PatchRunsIDReader{formats: a.formats},
		AuthInfo:           authInfo,
		Context:            params.Context,
		Client:             params.HTTPClient,
	})
	if err != nil {
		return nil, err
	}
	return result.(*PatchRunsIDOK), nil

}

/*
PatchRunsIDLogs links a log to a specific run
*/
func (a *Client) PatchRunsIDLogs(params *PatchRunsIDLogsParams, authInfo runtime.ClientAuthInfoWriter) (*PatchRunsIDLogsNoContent, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewPatchRunsIDLogsParams()
	}

	result, err := a.transport.Submit(&runtime.ClientOperation{
		ID:                 "PatchRunsIDLogs",
		Method:             "PATCH",
		PathPattern:        "/runs/{id}/logs",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"http"},
		Params:             params,
		Reader:             &PatchRunsIDLogsReader{formats: a.formats},
		AuthInfo:           authInfo,
		Context:            params.Context,
		Client:             params.HTTPClient,
	})
	if err != nil {
		return nil, err
	}
	return result.(*PatchRunsIDLogsNoContent), nil

}

/*
PostRuns creates a run
*/
func (a *Client) PostRuns(params *PostRunsParams, authInfo runtime.ClientAuthInfoWriter) (*PostRunsCreated, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewPostRunsParams()
	}

	result, err := a.transport.Submit(&runtime.ClientOperation{
		ID:                 "PostRuns",
		Method:             "POST",
		PathPattern:        "/runs",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"http"},
		Params:             params,
		Reader:             &PostRunsReader{formats: a.formats},
		AuthInfo:           authInfo,
		Context:            params.Context,
		Client:             params.HTTPClient,
	})
	if err != nil {
		return nil, err
	}
	return result.(*PostRunsCreated), nil

}

// SetTransport changes the transport on the client
func (a *Client) SetTransport(transport runtime.ClientTransport) {
	a.transport = transport
}
