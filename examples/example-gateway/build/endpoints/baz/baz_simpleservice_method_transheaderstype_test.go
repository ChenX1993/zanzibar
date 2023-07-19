// Code generated by zanzibar
// @generated

// Copyright (c) 2018 Uber Technologies, Inc.
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

package bazendpoint

import (
	"bytes"
	"context"
	"encoding/json"
	"io"
	"strconv"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/uber/zanzibar/v2/config"
	testbackend "github.com/uber/zanzibar/v2/test/lib/test_backend"
	testGateway "github.com/uber/zanzibar/v2/test/lib/test_gateway"
	"github.com/uber/zanzibar/v2/test/lib/util"

	bazclient "github.com/uber/zanzibar/v2/examples/example-gateway/build/clients/baz"
	clientsIDlClientsBazBaz "github.com/uber/zanzibar/v2/examples/example-gateway/build/gen-code/clients-idl/clients/baz/baz"
)

func TestTransHeadersTypeSuccessfulRequestOKResponse(t *testing.T) {

	confFiles := util.DefaultConfigFiles("example-gateway")
	staticConf := config.NewRuntimeConfigOrDie(confFiles, map[string]interface{}{})
	var alternateServiceDetail config.AlternateServiceDetail
	if staticConf.ContainsKey("clients.baz.alternates") {
		staticConf.MustGetStruct("clients.baz.alternates", &alternateServiceDetail)
	}
	var backends []*testbackend.TestTChannelBackend
	for serviceName := range alternateServiceDetail.ServicesDetailMap {
		if serviceName == "nomatch" {
			continue
		}
		backend, err := testbackend.CreateTChannelBackend(int32(0), serviceName)
		assert.NoError(t, err)
		err = backend.Bootstrap()
		assert.NoError(t, err)
		backends = append(backends, backend)
	}

	gateway, err := testGateway.CreateGateway(t, map[string]interface{}{
		"clients.baz.serviceName": "bazService",
	}, &testGateway.Options{
		KnownTChannelBackends: []string{"baz"},
		TestBinary:            util.DefaultMainFile("example-gateway"),
		ConfigFiles:           confFiles,
		Backends:              backends,
	})
	if !assert.NoError(t, err, "got bootstrap err") {
		return
	}
	defer gateway.Close()

	fakeTransHeadersType := func(
		ctx context.Context,
		reqHeaders map[string]string,
		args *clientsIDlClientsBazBaz.SimpleService_TransHeadersType_Args,
	) (*clientsIDlClientsBazBaz.TransHeaderType, map[string]string, error) {

		var resHeaders map[string]string

		var res clientsIDlClientsBazBaz.TransHeaderType

		clientResponse := []byte(`{"b1":true,"f3":3.14,"i1":3,"i2":3,"s6":"2a629c1a-262a-43f0-8623-869b0256a321","u4":"2a629c1a-262a-43f0-8623-869b0256a321","u5":"2a629c1a-262a-43f0-8623-869b0256a321"}`)
		err := json.Unmarshal(clientResponse, &res)
		if err != nil {
			t.Fatal("cant't unmarshal client response json to client response struct")
			return nil, resHeaders, err
		}
		return &res, resHeaders, nil
	}

	headers := map[string]string{}
	err = gateway.TChannelBackends()["baz"].Register(
		"baz", "transHeadersType", "SimpleService::transHeadersType",
		bazclient.NewSimpleServiceTransHeadersTypeHandler(fakeTransHeadersType),
	)
	assert.NoError(t, err)
	makeRequestAndValidateTransHeadersTypeSuccessfulRequest(t, gateway, headers)

	isSet := true
	i := 1
	for serviceName := range alternateServiceDetail.ServicesDetailMap {
		headers := map[string]string{}

		if serviceName == "nomatch" {
			headers["x-container"] = "randomstr"
			headers["x-test-Env"] = "randomstr"
		} else {
			if isSet {
				headers["x-container"] = "sandbox"
				isSet = false
			} else {
				headers["x-test-Env"] = "test1"
			}
			err = gateway.TChannelBackends()["baz:"+strconv.Itoa(i)].Register(
				"baz", "transHeadersType", "SimpleService::transHeadersType",
				bazclient.NewSimpleServiceTransHeadersTypeHandler(fakeTransHeadersType),
			)
			assert.NoError(t, err)
			i++
		}

		makeRequestAndValidateTransHeadersTypeSuccessfulRequest(t, gateway, headers)
	}

}

func makeRequestAndValidateTransHeadersTypeSuccessfulRequest(t *testing.T, gateway testGateway.TestGateway, headers map[string]string) {
	headers["x-boolean"] = "true"
	headers["x-float"] = "3.14"
	headers["x-int"] = "3"
	headers["x-string"] = "2a629c1a-262a-43f0-8623-869b0256a321"

	endpointRequest := []byte(`{"req":{}}`)

	res, err := gateway.MakeRequest(
		"POST",
		"/baz/trans-header-type",
		headers,
		bytes.NewReader(endpointRequest),
	)
	if !assert.NoError(t, err, "got http error") {
		return
	}

	defer func() { _ = res.Body.Close() }()
	data, err := io.ReadAll(res.Body)
	if !assert.NoError(t, err, "failed to read response body") {
		return
	}

	assert.Equal(t, 200, res.StatusCode)
	assert.JSONEq(t, `{}`, string(data))
}
