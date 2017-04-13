// Code generated by zanzibar
// @generated

package baz

import (
	"context"

	"github.com/uber/zanzibar/examples/example-gateway/build/clients"
	zanzibar "github.com/uber/zanzibar/runtime"
	"go.uber.org/zap"

	endpointsBazBaz "github.com/uber/zanzibar/examples/example-gateway/build/gen-code/endpoints/baz/baz"
	customBaz "github.com/uber/zanzibar/examples/example-gateway/endpoints/baz"
)

// HandleCallRequest handles "/baz/call-path".
func HandleCallRequest(
	ctx context.Context,
	req *zanzibar.ServerHTTPRequest,
	res *zanzibar.ServerHTTPResponse,
	clients *clients.Clients,
) {
	var requestBody endpointsBazBaz.BazRequest
	if ok := req.ReadAndUnmarshalBody(&requestBody); !ok {
		return
	}

	headers := map[string]string{}
	// TODO(sindelar): Add optional headers in addition to required.
	for _, k := range []string(nil) {
		headers[k] = req.Header.Get(k)
	}

	workflow := customBaz.CallEndpoint{
		Clients: clients,
		Logger:  req.Logger,
		Request: req,
	}

	response, respHeaders, err := workflow.Handle(ctx, headers, &requestBody)
	if err != nil {
		req.Logger.Warn("Workflow for endpoint returned error",
			zap.String("error", err.Error()),
		)
		res.SendErrorString(500, "Unexpected server error")
		return
	}

	// TODO(sindelar): Add response headers as an thrift spec annotation.
	endRespHead := map[string]string{}
	for _, k := range []string(nil) {
		endRespHead[k] = respHeaders[k]
	}

	res.WriteJSON(200, endRespHead, response)
}
