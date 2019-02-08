// Code generated by go-swagger; DO NOT EDIT.

package client

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"github.com/go-openapi/runtime"
	httptransport "github.com/go-openapi/runtime/client"

	strfmt "github.com/go-openapi/strfmt"

	"github.com/SoftwareForScience/jiskefet-api-go/client/attachments"
	"github.com/SoftwareForScience/jiskefet-api-go/client/authentication"
	"github.com/SoftwareForScience/jiskefet-api-go/client/logs"
	"github.com/SoftwareForScience/jiskefet-api-go/client/overview"
	"github.com/SoftwareForScience/jiskefet-api-go/client/runs"
	"github.com/SoftwareForScience/jiskefet-api-go/client/subsystems"
	"github.com/SoftwareForScience/jiskefet-api-go/client/users"
)

// Default jiskefet HTTP client.
var Default = NewHTTPClient(nil)

const (
	// DefaultHost is the default Host
	// found in Meta (info) section of spec file
	DefaultHost string = "localhost"
	// DefaultBasePath is the default BasePath
	// found in Meta (info) section of spec file
	DefaultBasePath string = "/api"
)

// DefaultSchemes are the default schemes found in Meta (info) section of spec file
var DefaultSchemes = []string{"http"}

// NewHTTPClient creates a new jiskefet HTTP client.
func NewHTTPClient(formats strfmt.Registry) *Jiskefet {
	return NewHTTPClientWithConfig(formats, nil)
}

// NewHTTPClientWithConfig creates a new jiskefet HTTP client,
// using a customizable transport config.
func NewHTTPClientWithConfig(formats strfmt.Registry, cfg *TransportConfig) *Jiskefet {
	// ensure nullable parameters have default
	if cfg == nil {
		cfg = DefaultTransportConfig()
	}

	// create transport and client
	transport := httptransport.New(cfg.Host, cfg.BasePath, cfg.Schemes)
	return New(transport, formats)
}

// New creates a new jiskefet client
func New(transport runtime.ClientTransport, formats strfmt.Registry) *Jiskefet {
	// ensure nullable parameters have default
	if formats == nil {
		formats = strfmt.Default
	}

	cli := new(Jiskefet)
	cli.Transport = transport

	cli.Attachments = attachments.New(transport, formats)

	cli.Authentication = authentication.New(transport, formats)

	cli.Logs = logs.New(transport, formats)

	cli.Overview = overview.New(transport, formats)

	cli.Runs = runs.New(transport, formats)

	cli.Subsystems = subsystems.New(transport, formats)

	cli.Users = users.New(transport, formats)

	return cli
}

// DefaultTransportConfig creates a TransportConfig with the
// default settings taken from the meta section of the spec file.
func DefaultTransportConfig() *TransportConfig {
	return &TransportConfig{
		Host:     DefaultHost,
		BasePath: DefaultBasePath,
		Schemes:  DefaultSchemes,
	}
}

// TransportConfig contains the transport related info,
// found in the meta section of the spec file.
type TransportConfig struct {
	Host     string
	BasePath string
	Schemes  []string
}

// WithHost overrides the default host,
// provided by the meta section of the spec file.
func (cfg *TransportConfig) WithHost(host string) *TransportConfig {
	cfg.Host = host
	return cfg
}

// WithBasePath overrides the default basePath,
// provided by the meta section of the spec file.
func (cfg *TransportConfig) WithBasePath(basePath string) *TransportConfig {
	cfg.BasePath = basePath
	return cfg
}

// WithSchemes overrides the default schemes,
// provided by the meta section of the spec file.
func (cfg *TransportConfig) WithSchemes(schemes []string) *TransportConfig {
	cfg.Schemes = schemes
	return cfg
}

// Jiskefet is a client for jiskefet
type Jiskefet struct {
	Attachments *attachments.Client

	Authentication *authentication.Client

	Logs *logs.Client

	Overview *overview.Client

	Runs *runs.Client

	Subsystems *subsystems.Client

	Users *users.Client

	Transport runtime.ClientTransport
}

// SetTransport changes the transport on the client and all its subresources
func (c *Jiskefet) SetTransport(transport runtime.ClientTransport) {
	c.Transport = transport

	c.Attachments.SetTransport(transport)

	c.Authentication.SetTransport(transport)

	c.Logs.SetTransport(transport)

	c.Overview.SetTransport(transport)

	c.Runs.SetTransport(transport)

	c.Subsystems.SetTransport(transport)

	c.Users.SetTransport(transport)

}
