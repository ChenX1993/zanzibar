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

package main

import (
	"flag"
	"os"
	"os/signal"
	"runtime"
	"strings"
	"syscall"

	"go.uber.org/zap"

	"github.com/uber/zanzibar/config"
	"github.com/uber/zanzibar/runtime"

	app "github.com/uber/zanzibar/examples/example-gateway"
	service "github.com/uber/zanzibar/examples/example-gateway/build/services/example-gateway"
)

var configFiles *string

func getDirName() string {
	_, file, _, _ := runtime.Caller(0)
	return zanzibar.GetDirnameFromRuntimeCaller(file)
}

func getConfig() *zanzibar.StaticConfig {
	var files []string

	if configFiles == nil {
		files = []string{}
	} else {
		files = strings.Split(*configFiles, ";")
	}

	return config.NewRuntimeConfigOrDie(files, nil)
}

func createGateway() (*zanzibar.Gateway, error) {
	config := getConfig()

	gateway, _, err := service.CreateGateway(config, app.AppOptions)
	if err != nil {
		return nil, err
	}

	return gateway, nil
}

func logAndWait(server *zanzibar.Gateway) {
	server.Logger.Info("Started ExampleGateway",
		zap.String("realHTTPAddr", server.RealHTTPAddr),
		zap.String("realTChannelAddr", server.RealTChannelAddr),
		zap.Any("config", server.InspectOrDie()),
	)

	go func() {
		sig := make(chan os.Signal, 1)
		signal.Notify(sig, syscall.SIGINT, syscall.SIGTERM)
		<-sig
		server.WaitGroup.Add(1)
		server.Shutdown()
		server.WaitGroup.Done()
	}()
	server.Wait()
}

func readFlags() {
	configFiles = flag.String(
		"config",
		"",
		"an ordered, semi-colon separated list of configuration files to use",
	)
	flag.Parse()
}

func main() {
	readFlags()
	server, err := createGateway()
	if err != nil {
		panic(err)
	}

	err = server.Bootstrap()
	if err != nil {
		panic(err)
	}

	logAndWait(server)
}
