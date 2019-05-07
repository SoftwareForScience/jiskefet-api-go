// Code generated by go-swagger; DO NOT EDIT.

package tags

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"github.com/go-openapi/runtime"

	strfmt "github.com/go-openapi/strfmt"
)

// New creates a new tags API client.
func New(transport runtime.ClientTransport, formats strfmt.Registry) *Client {
	return &Client{transport: transport, formats: formats}
}

/*
Client for tags API
*/
type Client struct {
	transport runtime.ClientTransport
	formats   strfmt.Registry
}

/*
DeleteTagsID deletes a specific tag
*/
func (a *Client) DeleteTagsID(params *DeleteTagsIDParams, authInfo runtime.ClientAuthInfoWriter) (*DeleteTagsIDNoContent, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewDeleteTagsIDParams()
	}

	result, err := a.transport.Submit(&runtime.ClientOperation{
		ID:                 "DeleteTagsID",
		Method:             "DELETE",
		PathPattern:        "/tags/{id}",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"http"},
		Params:             params,
		Reader:             &DeleteTagsIDReader{formats: a.formats},
		AuthInfo:           authInfo,
		Context:            params.Context,
		Client:             params.HTTPClient,
	})
	if err != nil {
		return nil, err
	}
	return result.(*DeleteTagsIDNoContent), nil

}

/*
GetTags returns all tags
*/
func (a *Client) GetTags(params *GetTagsParams, authInfo runtime.ClientAuthInfoWriter) (*GetTagsOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewGetTagsParams()
	}

	result, err := a.transport.Submit(&runtime.ClientOperation{
		ID:                 "GetTags",
		Method:             "GET",
		PathPattern:        "/tags",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"http"},
		Params:             params,
		Reader:             &GetTagsReader{formats: a.formats},
		AuthInfo:           authInfo,
		Context:            params.Context,
		Client:             params.HTTPClient,
	})
	if err != nil {
		return nil, err
	}
	return result.(*GetTagsOK), nil

}

/*
GetTagsID returns a specific tag
*/
func (a *Client) GetTagsID(params *GetTagsIDParams, authInfo runtime.ClientAuthInfoWriter) (*GetTagsIDOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewGetTagsIDParams()
	}

	result, err := a.transport.Submit(&runtime.ClientOperation{
		ID:                 "GetTagsID",
		Method:             "GET",
		PathPattern:        "/tags/{id}",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"http"},
		Params:             params,
		Reader:             &GetTagsIDReader{formats: a.formats},
		AuthInfo:           authInfo,
		Context:            params.Context,
		Client:             params.HTTPClient,
	})
	if err != nil {
		return nil, err
	}
	return result.(*GetTagsIDOK), nil

}

/*
GetTagsIDLogs returns all logs for a specific tag
*/
func (a *Client) GetTagsIDLogs(params *GetTagsIDLogsParams, authInfo runtime.ClientAuthInfoWriter) (*GetTagsIDLogsOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewGetTagsIDLogsParams()
	}

	result, err := a.transport.Submit(&runtime.ClientOperation{
		ID:                 "GetTagsIDLogs",
		Method:             "GET",
		PathPattern:        "/tags/{id}/logs",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"http"},
		Params:             params,
		Reader:             &GetTagsIDLogsReader{formats: a.formats},
		AuthInfo:           authInfo,
		Context:            params.Context,
		Client:             params.HTTPClient,
	})
	if err != nil {
		return nil, err
	}
	return result.(*GetTagsIDLogsOK), nil

}

/*
GetTagsIDRuns returns all runs for a specific tag
*/
func (a *Client) GetTagsIDRuns(params *GetTagsIDRunsParams, authInfo runtime.ClientAuthInfoWriter) (*GetTagsIDRunsOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewGetTagsIDRunsParams()
	}

	result, err := a.transport.Submit(&runtime.ClientOperation{
		ID:                 "GetTagsIDRuns",
		Method:             "GET",
		PathPattern:        "/tags/{id}/runs",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"http"},
		Params:             params,
		Reader:             &GetTagsIDRunsReader{formats: a.formats},
		AuthInfo:           authInfo,
		Context:            params.Context,
		Client:             params.HTTPClient,
	})
	if err != nil {
		return nil, err
	}
	return result.(*GetTagsIDRunsOK), nil

}

/*
PatchTagsIDLogs links a run to a specific tag
*/
func (a *Client) PatchTagsIDLogs(params *PatchTagsIDLogsParams, authInfo runtime.ClientAuthInfoWriter) (*PatchTagsIDLogsNoContent, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewPatchTagsIDLogsParams()
	}

	result, err := a.transport.Submit(&runtime.ClientOperation{
		ID:                 "PatchTagsIDLogs",
		Method:             "PATCH",
		PathPattern:        "/tags/{id}/logs",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"http"},
		Params:             params,
		Reader:             &PatchTagsIDLogsReader{formats: a.formats},
		AuthInfo:           authInfo,
		Context:            params.Context,
		Client:             params.HTTPClient,
	})
	if err != nil {
		return nil, err
	}
	return result.(*PatchTagsIDLogsNoContent), nil

}

/*
PatchTagsIDRuns links a run to a specific tag
*/
func (a *Client) PatchTagsIDRuns(params *PatchTagsIDRunsParams, authInfo runtime.ClientAuthInfoWriter) (*PatchTagsIDRunsNoContent, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewPatchTagsIDRunsParams()
	}

	result, err := a.transport.Submit(&runtime.ClientOperation{
		ID:                 "PatchTagsIDRuns",
		Method:             "PATCH",
		PathPattern:        "/tags/{id}/runs",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"http"},
		Params:             params,
		Reader:             &PatchTagsIDRunsReader{formats: a.formats},
		AuthInfo:           authInfo,
		Context:            params.Context,
		Client:             params.HTTPClient,
	})
	if err != nil {
		return nil, err
	}
	return result.(*PatchTagsIDRunsNoContent), nil

}

/*
PostTags creates a tag
*/
func (a *Client) PostTags(params *PostTagsParams, authInfo runtime.ClientAuthInfoWriter) (*PostTagsCreated, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewPostTagsParams()
	}

	result, err := a.transport.Submit(&runtime.ClientOperation{
		ID:                 "PostTags",
		Method:             "POST",
		PathPattern:        "/tags",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"http"},
		Params:             params,
		Reader:             &PostTagsReader{formats: a.formats},
		AuthInfo:           authInfo,
		Context:            params.Context,
		Client:             params.HTTPClient,
	})
	if err != nil {
		return nil, err
	}
	return result.(*PostTagsCreated), nil

}

// SetTransport changes the transport on the client
func (a *Client) SetTransport(transport runtime.ClientTransport) {
	a.transport = transport
}
