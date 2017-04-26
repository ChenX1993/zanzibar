// Code generated by zanzibar
// @generated

// Copyright (c) 2017 Uber Technologies, Inc.
//
// Permission is hereby granted, free of charge, to any person obtaining a copy
// of this software and associated documentation files (the "Software"), to deal
// in the Software without restriction, including without limitation the rights
// to use, copy, modify, merge, publish, distribute, sublicense, and/or sell
// copies of the Software, and to permit persons to whom the Software is
// furnished to do so, subject to the following conditions:
//
// The above copyright notice and this permission notice shall be included in
// all copies or substantial portions of the Software.
//
// THE SOFTWARE IS PROVIDED "AS IS", WITHOUT WARRANTY OF ANY KIND, EXPRESS OR
// IMPLIED, INCLUDING BUT NOT LIMITED TO THE WARRANTIES OF MERCHANTABILITY,
// FITNESS FOR A PARTICULAR PURPOSE AND NONINFRINGEMENT. IN NO EVENT SHALL THE
// AUTHORS OR COPYRIGHT HOLDERS BE LIABLE FOR ANY CLAIM, DAMAGES OR OTHER
// LIABILITY, WHETHER IN AN ACTION OF CONTRACT, TORT OR OTHERWISE, ARISING FROM,
// OUT OF OR IN CONNECTION WITH THE SOFTWARE OR THE USE OR OTHER DEALINGS IN
// THE SOFTWARE.

package baz

import (
	"context"

	"github.com/uber/zanzibar/examples/example-gateway/build/clients"
	zanzibar "github.com/uber/zanzibar/runtime"
	"go.uber.org/zap"

	clientsBazBaz "github.com/uber/zanzibar/examples/example-gateway/build/gen-code/clients/baz/baz"
	endpointsBazBaz "github.com/uber/zanzibar/examples/example-gateway/build/gen-code/endpoints/baz/baz"
)

// SillyNoopHandler is the handler for "/baz/silly-noop"
type SillyNoopHandler struct {
	Clients *clients.Clients
}

// NewSillyNoopEndpoint creates a handler
func NewSillyNoopEndpoint(
	gateway *zanzibar.Gateway,
) *SillyNoopHandler {
	return &SillyNoopHandler{
		Clients: gateway.Clients.(*clients.Clients),
	}
}

// HandleRequest handles "/baz/silly-noop".
func (handler *SillyNoopHandler) HandleRequest(
	ctx context.Context,
	req *zanzibar.ServerHTTPRequest,
	res *zanzibar.ServerHTTPResponse,
) {

	workflow := SillyNoopEndpoint{
		Clients: handler.Clients,
		Logger:  req.Logger,
		Request: req,
	}

	cliRespHeaders, err := workflow.Handle(ctx, req.Header)
	if err != nil {
		switch errValue := err.(type) {

		case *endpointsBazBaz.AuthErr:
			res.WriteJSON(
				403, cliRespHeaders, errValue,
			)
			return

		case *endpointsBazBaz.ServerErr:
			res.WriteJSON(
				500, cliRespHeaders, errValue,
			)
			return

		default:
			req.Logger.Warn("Workflow for endpoint returned error",
				zap.String("error", errValue.Error()),
			)
			res.SendErrorString(500, "Unexpected server error")
			return
		}
	}

	res.WriteJSONBytes(204, cliRespHeaders, nil)
}

// SillyNoopEndpoint calls thrift client Baz.SillyNoop
type SillyNoopEndpoint struct {
	Clients *clients.Clients
	Logger  *zap.Logger
	Request *zanzibar.ServerHTTPRequest
}

// Handle calls thrift client.
func (w SillyNoopEndpoint) Handle(
	ctx context.Context,
	reqHeaders zanzibar.Header,
) (zanzibar.Header, error) {

	clientHeaders := map[string]string{}

	_, err := w.Clients.Baz.SillyNoop(ctx, clientHeaders)

	if err != nil {
		switch errValue := err.(type) {

		case *clientsBazBaz.AuthErr:
			serverErr := convertSillyNoopAuthErr(
				errValue,
			)
			// TODO(sindelar): Consider returning partial headers

			return nil, serverErr

		case *clientsBazBaz.ServerErr:
			serverErr := convertSillyNoopServerErr(
				errValue,
			)
			// TODO(sindelar): Consider returning partial headers

			return nil, serverErr

		default:
			w.Logger.Warn("Could not make client request",
				zap.String("error", errValue.Error()),
			)
			// TODO(sindelar): Consider returning partial headers

			return nil, err

		}
	}

	// Filter and map response headers from client to server response.

	// TODO: Add support for TChannel Headers with a switch here
	resHeaders := zanzibar.ServerHTTPHeader{}

	return resHeaders, nil
}

func convertSillyNoopAuthErr(
	clientError *clientsBazBaz.AuthErr,
) *endpointsBazBaz.AuthErr {
	// TODO: Add error fields mapping here.
	serverError := &endpointsBazBaz.AuthErr{}
	return serverError
}
func convertSillyNoopServerErr(
	clientError *clientsBazBaz.ServerErr,
) *endpointsBazBaz.ServerErr {
	// TODO: Add error fields mapping here.
	serverError := &endpointsBazBaz.ServerErr{}
	return serverError
}