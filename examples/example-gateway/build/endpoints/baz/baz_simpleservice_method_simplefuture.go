// Code generated by zanzibar
// @generated

package baz

import (
	"context"

	"github.com/uber/zanzibar/examples/example-gateway/build/clients"
	zanzibar "github.com/uber/zanzibar/runtime"
	"go.uber.org/zap"

	customBaz "github.com/uber/zanzibar/examples/example-gateway/endpoints/baz"
)

// HandleSimpleFutureRequest handles "/baz/simple-future-path".
func HandleSimpleFutureRequest(
	ctx context.Context,
	req *zanzibar.ServerHTTPRequest,
	res *zanzibar.ServerHTTPResponse,
	clients *clients.Clients,
) {

	headers := map[string]string{}
	// TODO(sindelar): Add optional headers in addition to required.
	for _, k := range []string(nil) {
		headers[k] = req.Header.Get(k)
	}

	workflow := customBaz.SimpleFutureEndpoint{
		Clients: clients,
		Logger:  req.Logger,
		Request: req,
	}

	respHeaders, err := workflow.Handle(ctx, headers)
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

	res.WriteJSONBytes(204, endRespHead, nil)
}
