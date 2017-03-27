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

package codegen

import (
	"encoding/json"
	"go/parser"
	"go/token"
	"io/ioutil"
	"os"
	"path/filepath"
	"runtime"
	"strings"

	"fmt"

	"github.com/pkg/errors"
	"github.com/uber/zanzibar/runtime"
)

func getDirName() string {
	_, file, _, _ := runtime.Caller(0)
	return zanzibar.GetDirnameFromRuntimeCaller(file)
}

var templateDir = filepath.Join(getDirName(), "templates", "*.tmpl")
var mandatoryClientFields = []string{
	"clientType",
	"clientId",
	"thriftFile",
	"thriftFileSha",
	"clientName",
	"serviceName",
}
var mandatoryEndpointFields = []string{
	"endpointType",
	"endpointId",
	"handleId",
	"thriftFile",
	"thriftFileSha",
	"thriftMethodName",
	"workflowType",
	"testFixtures",
	"middlewares",
}

// ClientSpec holds information about each client in the
// gateway included its thriftFile and other meta info
type ClientSpec struct {
	// ModuleSpec holds the thrift module information
	ModuleSpec *ModuleSpec
	// JSONFile for this spec
	JSONFile string
	// ClientType, currently only "http" is supported
	ClientType string
	// GoFileName, the absolute path where the generate client is
	GoFileName string
	// GoPackageName is the golang package name for the client
	GoPackageName string
	// GoStructsFileName, absolute path where any helper structs
	// are generated for this generated client
	GoStructsFileName string
	// ThriftFile, absolute path to thrift file
	ThriftFile string
	// ClientID, used for logging and metrics, must be lowercase
	// and use dashes.
	ClientID string
	// ClientName, PascalCase name of the client, the generated
	// `Clients` struct will contain a field of this name
	ClientName string
	// ThriftServiceName, if the thrift file has multiple
	// services then this is the service that describes the client
	ThriftServiceName string
}

// NewClientSpec creates a client spec from a json file.
func NewClientSpec(jsonFile string, h *PackageHelper) (*ClientSpec, error) {
	_, err := os.Stat(jsonFile)
	if err != nil {
		return nil, errors.Wrapf(err, "Could not find file %s: ", jsonFile)
	}

	bytes, err := ioutil.ReadFile(jsonFile)
	if err != nil {
		return nil, errors.Wrapf(
			err, "Could not read json file %s: ", jsonFile,
		)
	}

	clientConfigGeneric := map[string]interface{}{}
	err = json.Unmarshal(bytes, &clientConfigGeneric)
	if err != nil {
		return nil, errors.Wrapf(
			err, "Could not parse json file %s: ", jsonFile,
		)
	}

	clientType := clientConfigGeneric["clientType"].(string)
	if clientType == "http" {
		// A better solution for client configs is needed here.
		// Assuming a string map is too restrictive
		clientConfigObj := map[string]string{}
		err := json.Unmarshal(bytes, &clientConfigObj)

		if err != nil {
			return nil, errors.Wrapf(
				err, "Could not parse json file %s: ", jsonFile,
			)
		}

		return NewHTTPClientSpec(jsonFile, clientConfigObj, h)
	} else if clientType == "tchannel" {
		return NewTChannelClientSpec(jsonFile, clientConfigGeneric, h)
	} else if clientType == "custom" {
		return NewCustomClientSpec(jsonFile, clientConfigGeneric, h)
	}

	return nil, errors.Errorf(
		"Cannot support unknown clientType for client %s", jsonFile,
	)
}

// NewTChannelClientSpec creates a client spec from a json file whose type is tchannel
func NewTChannelClientSpec(jsonFile string, clientConfigObj map[string]interface{}, h *PackageHelper) (*ClientSpec, error) {
	return &ClientSpec{
		ClientType: "tchannel",
	}, nil
}

// NewCustomClientSpec creates a client spec from a json file whose type is custom
func NewCustomClientSpec(jsonFile string, clientConfigObj map[string]interface{}, h *PackageHelper) (*ClientSpec, error) {
	return &ClientSpec{
		ClientType: "custom",
	}, nil
}

// NewHTTPClientSpec creates a client spec from a json file whose type is http
func NewHTTPClientSpec(jsonFile string, clientConfigObj map[string]string, h *PackageHelper) (*ClientSpec, error) {
	for i := 0; i < len(mandatoryClientFields); i++ {
		fieldName := mandatoryClientFields[i]
		if clientConfigObj[fieldName] == "" {
			return nil, errors.Errorf(
				"client config (%s) must have %s field", jsonFile, fieldName,
			)
		}
	}

	thriftFile := filepath.Join(
		h.ThriftIDLPath(), clientConfigObj["thriftFile"],
	)

	mspec, err := NewModuleSpec(thriftFile, h)
	if err != nil {
		return nil, errors.Wrapf(
			err, "Could not build module spec for thrift %s: ", thriftFile,
		)
	}

	baseName := filepath.Base(filepath.Dir(jsonFile))

	goFileName := filepath.Join(
		h.CodeGenTargetPath(),
		"clients",
		baseName,
		baseName+".go",
	)

	goPackageName := filepath.Join(
		h.GoGatewayPackageName(),
		"clients",
		baseName,
	)

	goStructsFileName := filepath.Join(
		h.CodeGenTargetPath(),
		"clients",
		baseName,
		baseName+"_structs.go",
	)

	return &ClientSpec{
		ModuleSpec:        mspec,
		JSONFile:          jsonFile,
		ClientType:        "http",
		GoFileName:        goFileName,
		GoPackageName:     goPackageName,
		GoStructsFileName: goStructsFileName,
		ThriftFile:        thriftFile,
		ClientID:          clientConfigObj["clientId"],
		ClientName:        clientConfigObj["clientName"],
		ThriftServiceName: clientConfigObj["serviceName"],
	}, nil
}

// MiddlewareSpec holds information about each middleware at the endpoint
// level. The same mid
type MiddlewareSpec struct {
	// The middleware package name.
	Name string
	// Go import path for the middleware.
	Path string
	// Middleware specific configuration options.
	Options map[string]interface{}
}

func NewMiddlewareSpec(goFile string, h *PackageHelper) (*MiddlewareSpec, error) {
	_, err := os.Stat(goFile)
	if err != nil {
		return nil, errors.Wrapf(err, "Could not find file %s: ", goFile)
	}

	// Read the go middleware definitions.
	fset := token.NewFileSet()
	f, err := parser.ParseFile(fset, goFile, nil, parser.PackageClauseOnly)
	if err != nil {
		return nil, errors.Wrapf(
			err, "Could not read go file %s: ", goFile,
		)
	}

	fmt.Printf("SINDELARTEST %s", string(f.Package))

	// TODO(sindelar): Consider adding other middleware validation here.
	return &MiddlewareSpec{
		Name: string(f.Package),
		Path: goFile,
	}, nil
}

// EndpointSpec holds information about each endpoint in the
// gateway including its thriftFile and meta data
type EndpointSpec struct {
	// ModuleSpec holds the thrift module info
	ModuleSpec *ModuleSpec
	// JSONFile for this endpoint spec
	JSONFile string
	// GoStructsFileName is where structs are generated
	GoStructsFileName string
	// GoFolderName is the folder where all the endpoints
	// are generated.
	GoFolderName string

	// EndpointType, currently only "http"
	EndpointType string
	// EndpointID, used in metrics and logging, lower case.
	EndpointID string
	// HandleID, used in metrics and logging, lowercase.
	HandleID string
	// ThriftFile, the thrift file for this endpoint
	ThriftFile string
	// ThriftMethodName, which thrift method to use.
	ThriftMethodName string
	// ThriftServiceName, which thrift service to use.
	ThriftServiceName string
	// TestFixtures, meta data to generate tests,
	// TODO figure out struct type
	TestFixtures []interface{}
	// Middlewares, meta data to add middlewares,
	Middlewares []MiddlewareSpec

	// WorkflowType, either "httpClient" or "custom".
	// A httpClient workflow generates a http client Caller
	// A custom workflow just imports the custom code
	WorkflowType string
	// If "custom" then where to import custom code from
	WorkflowImportPath string
	// if "httpClient", which client to call.
	ClientName string
	// if "httpClient", which client method to call.
	ClientMethod string
}

// NewEndpointSpec creats an endpoint spec from a json file.
func NewEndpointSpec(
	jsonFile string,
	h *PackageHelper,
	midSpecs map[string]*MiddlewareSpec,
) (*EndpointSpec, error) {
	_, err := os.Stat(jsonFile)
	if err != nil {
		return nil, errors.Wrapf(err, "Could not find file %s: ", jsonFile)
	}

	bytes, err := ioutil.ReadFile(jsonFile)
	if err != nil {
		return nil, errors.Wrapf(
			err, "Could not read json file %s: ", jsonFile,
		)
	}

	endpointConfigObj := map[string]interface{}{}
	err = json.Unmarshal(bytes, &endpointConfigObj)
	if err != nil {
		return nil, errors.Wrapf(
			err, "Could not parse json file %s: ", jsonFile,
		)
	}

	if endpointConfigObj["endpointType"] != "http" {
		return nil, errors.Errorf(
			"Cannot support unknown endpointType for endpoint %s", jsonFile,
		)
	}

	for i := 0; i < len(mandatoryEndpointFields); i++ {
		fieldName := mandatoryEndpointFields[i]
		if endpointConfigObj[fieldName] == "" {
			return nil, errors.Errorf(
				"endpoint config (%s) must have %s field", jsonFile, fieldName,
			)
		}
	}

	thriftFile := filepath.Join(
		h.ThriftIDLPath(), endpointConfigObj["thriftFile"].(string),
	)

	mspec, err := NewModuleSpec(thriftFile, h)
	if err != nil {
		return nil, errors.Wrapf(
			err, "Could not build module spec for thrift %s: ", thriftFile,
		)
	}

	var workflowImportPath string
	var clientName string
	var clientMethod string

	workflowType := endpointConfigObj["workflowType"].(string)
	if workflowType == "httpClient" {
		iclientName, ok := endpointConfigObj["clientName"]
		if !ok {
			return nil, errors.Errorf(
				"endpoint config (%s) must have clientName field", jsonFile,
			)
		}
		clientName = iclientName.(string)

		iclientMethod, ok := endpointConfigObj["clientMethod"]
		if !ok {
			return nil, errors.Errorf(
				"endpoint config (%s) must have clientMethod field", jsonFile,
			)
		}
		clientMethod = iclientMethod.(string)
	} else if workflowType == "custom" {
		iworkflowImportPath, ok := endpointConfigObj["workflowImportPath"]
		if !ok {
			return nil, errors.Errorf(
				"endpoint config (%s) must have workflowImportPath field",
				jsonFile,
			)
		}
		workflowImportPath = iworkflowImportPath.(string)
	} else {
		return nil, errors.Errorf(
			"Invalid workflowType (%s) for endpoint (%s)",
			workflowType, jsonFile,
		)
	}

	dirName := filepath.Base(filepath.Dir(jsonFile))

	goFolderName := filepath.Join(
		h.CodeGenTargetPath(),
		"endpoints",
		dirName,
	)

	goStructsFileName := filepath.Join(
		h.CodeGenTargetPath(),
		"endpoints",
		dirName,
		dirName+"_structs.go",
	)

	thriftInfo := endpointConfigObj["thriftMethodName"].(string)
	parts := strings.Split(thriftInfo, "::")
	if len(parts) != 2 {
		return nil, errors.Errorf(
			"Cannot read thriftMethodName (%s) for endpoint json file %s : ",
			thriftInfo, jsonFile,
		)
	}

	// TODO(sindelar): Validate middleware are defined in codebase.
	middlewares := make([]MiddlewareSpec, len(endpointConfigObj["middlewares"].([]interface{})))
	for idx, middleware := range endpointConfigObj["middlewares"].([]interface{}) {
		middlewareObj := middleware.(map[string]interface{})
		name := middlewareObj["name"].(string)
		// Verify the middleware name is defined.
		if midSpecs[name] == nil {
			return nil, errors.Errorf(
				"middlewares config (%s) not found.", name,
			)
		}
		// TODO(sindelar): Validate Options against middleware spec and support
		// nested typed objects.
		middlewares[idx] = MiddlewareSpec{
			Name:    name,
			Path:    midSpecs[name].Path,
			Options: middlewareObj["options"].(map[string]interface{}),
		}
	}

	return &EndpointSpec{
		ModuleSpec:         mspec,
		JSONFile:           jsonFile,
		GoStructsFileName:  goStructsFileName,
		GoFolderName:       goFolderName,
		EndpointType:       endpointConfigObj["endpointType"].(string),
		EndpointID:         endpointConfigObj["endpointId"].(string),
		HandleID:           endpointConfigObj["handleId"].(string),
		ThriftFile:         thriftFile,
		ThriftServiceName:  parts[0],
		ThriftMethodName:   parts[1],
		TestFixtures:       endpointConfigObj["testFixtures"].([]interface{}),
		Middlewares:        middlewares,
		WorkflowType:       workflowType,
		WorkflowImportPath: workflowImportPath,
		ClientName:         clientName,
		ClientMethod:       clientMethod,
	}, nil
}

// TargetEndpointPath generates a filepath for each endpoint method
func (e *EndpointSpec) TargetEndpointPath(
	serviceName string, methodName string,
) string {
	baseName := filepath.Base(e.GoFolderName)

	fileName := baseName + "_" + strings.ToLower(serviceName) +
		"_method_" + strings.ToLower(methodName) + ".go"
	return filepath.Join(e.GoFolderName, fileName)
}

// TargetEndpointTestPath generates a filepath for each endpoint test
func (e *EndpointSpec) TargetEndpointTestPath(
	serviceName string, methodName string,
) string {
	baseName := filepath.Base(e.GoFolderName)

	fileName := baseName + "_" + strings.ToLower(serviceName) +
		"_method_" + strings.ToLower(methodName) + "_test.go"
	return filepath.Join(e.GoFolderName, fileName)
}

// EndpointTestConfigPath generates a filepath for each endpoint test config
func (e *EndpointSpec) EndpointTestConfigPath() string {
	return strings.TrimSuffix(e.JSONFile, filepath.Ext(e.JSONFile)) + "_test.json"
}

// SetDownstream configures the downstream client for this endpoint spec
func (e *EndpointSpec) SetDownstream(
	gatewaySpec *GatewaySpec,
) error {
	if e.WorkflowType == "custom" {
		return nil
	}

	var clientSpec *ClientSpec
	for _, v := range gatewaySpec.ClientModules {
		if v.ClientName == e.ClientName {
			clientSpec = v
			break
		}
	}

	if clientSpec == nil {
		return errors.Errorf(
			"When parsing endpoint json (%s), "+
				"could not find client (%s) in gateway",
			e.JSONFile, e.ClientName,
		)
	}

	return e.ModuleSpec.SetDownstream(
		e.ThriftServiceName, e.ThriftMethodName,
		clientSpec, e.ClientName, e.ClientMethod,
	)
}

// EndpointGroupInfo ...
type EndpointGroupInfo struct {
	Endpoints []string
}

func parseEndpointJsons(
	endpointGroupJsons []string,
) ([]string, error) {
	endpointJsons := []string{}

	for _, endpointGroupJSON := range endpointGroupJsons {
		bytes, err := ioutil.ReadFile(endpointGroupJSON)
		if err != nil {
			return nil, errors.Wrapf(
				err, "Cannot read endpoint group json: %s",
				endpointGroupJSON,
			)
		}

		var endpointGroupInfo EndpointGroupInfo
		err = json.Unmarshal(bytes, &endpointGroupInfo)
		if err != nil {
			return nil, errors.Wrapf(
				err, "Cannot parse json for endpoint group config: %s",
				endpointGroupJSON,
			)
		}

		endpointConfigDir := filepath.Dir(endpointGroupJSON)
		for _, fileName := range endpointGroupInfo.Endpoints {
			endpointJsons = append(
				endpointJsons, filepath.Join(endpointConfigDir, fileName),
			)
		}
	}

	return endpointJsons, nil
}

// GatewaySpec collects information for the entire gateway
type GatewaySpec struct {
	// package helper for gateway
	PackageHelper *PackageHelper
	// tempalte instance for gateway
	Template *Template

	ClientModules     map[string]*ClientSpec
	EndpointModules   map[string]*EndpointSpec
	MiddlewareModules map[string]*MiddlewareSpec

	gatewayName       string
	configDirName     string
	clientConfigDir   string
	endpointConfigDir string
	middlewareDir     string
}

// NewGatewaySpec sets up gateway spec
func NewGatewaySpec(
	configDirName string,
	thriftRootDir string,
	genCodePackage string,
	targetGenDir string,
	gatewayThriftRootDir string,
	clientConfig string,
	endpointConfig string,
	middlewareDir string,
	gatewayName string,
) (*GatewaySpec, error) {
	packageHelper, err := NewPackageHelper(
		thriftRootDir,
		genCodePackage,
		targetGenDir,
		gatewayThriftRootDir,
	)
	if err != nil {
		return nil, errors.Wrap(err, "Cannot build package helper")
	}

	tmpl, err := NewTemplate(templateDir)
	if err != nil {
		return nil, errors.Wrap(err, "cannot create template")
	}

	// TODO(sindelar): Filter test files.
	middlewareFiles, err := filepath.Glob(filepath.Join(
		configDirName,
		middlewareDir,
		"*",
		"*.go",
	))
	if err != nil {
		return nil, errors.Wrap(err, "Cannot load middlewares.")
	}

	clientJsons, err := filepath.Glob(filepath.Join(
		configDirName,
		clientConfig,
		"*",
		"client-config.json",
	))
	if err != nil {
		return nil, errors.Wrap(err, "Cannot load client json files")
	}

	endpointGroupJsons, err := filepath.Glob(filepath.Join(
		configDirName,
		endpointConfig,
		"*",
		"endpoint-config.json",
	))
	if err != nil {
		return nil, errors.Wrap(err, "Cannot load endpoint json files")
	}

	endpointJsons, err := parseEndpointJsons(endpointGroupJsons)
	if err != nil {
		return nil, errors.Wrap(err, "Cannot parse endpoint config")
	}

	spec := &GatewaySpec{
		PackageHelper:     packageHelper,
		Template:          tmpl,
		ClientModules:     map[string]*ClientSpec{},
		EndpointModules:   map[string]*EndpointSpec{},
		MiddlewareModules: map[string]*MiddlewareSpec{},

		configDirName:     configDirName,
		clientConfigDir:   clientConfig,
		endpointConfigDir: endpointConfig,
		gatewayName:       gatewayName,
	}
	for _, mid := range middlewareFiles {
		mspec, err := NewMiddlewareSpec(mid, packageHelper)
		if err != nil {
			return nil, errors.Wrapf(
				err, "Cannot parse middleware file %s :", mid,
			)
		}

		spec.MiddlewareModules[mspec.Name] = mspec
	}
	for _, json := range clientJsons {
		cspec, err := NewClientSpec(json, packageHelper)
		if err != nil {
			return nil, errors.Wrapf(
				err, "Cannot parse client json file %s :", json,
			)
		}

		// TODO: Other clients should be generated
		if cspec.ClientType == "http" {
			spec.ClientModules[cspec.JSONFile] = cspec
		}
	}
	for _, json := range endpointJsons {
		espec, err := NewEndpointSpec(json, packageHelper, spec.MiddlewareModules)
		if err != nil {
			return nil, errors.Wrapf(
				err, "Cannot parse endpoint json file %s :", json,
			)
		}

		err = espec.SetDownstream(spec)
		if err != nil {
			return nil, errors.Wrapf(
				err, "Cannot parse downstream info for endpoint : %s", json,
			)
		}
		spec.EndpointModules[espec.JSONFile] = espec
	}

	return spec, nil
}

// GenerateClients will generate all the clients for the gateway
func (gateway *GatewaySpec) GenerateClients() error {
	for _, module := range gateway.ClientModules {
		_, err := gateway.Template.GenerateClientFile(
			module, gateway.PackageHelper,
		)
		if err != nil {
			return err
		}
	}

	_, err := gateway.Template.GenerateClientsInitFile(
		gateway.ClientModules, gateway.PackageHelper,
	)
	return err
}

// GenerateEndpoints will generate all the endpoints for the gateway
func (gateway *GatewaySpec) GenerateEndpoints() error {
	for _, module := range gateway.EndpointModules {
		if module.WorkflowType == "custom" {
			continue
		}

		_, err := gateway.Template.GenerateEndpointFile(
			module, gateway.PackageHelper,
			module.ThriftServiceName, module.ThriftMethodName,
		)
		if err != nil {
			return err
		}

		_, err = gateway.Template.GenerateEndpointTestFile(
			module, gateway.PackageHelper,
			module.ThriftServiceName, module.ThriftMethodName,
		)
		if err != nil {
			return err
		}
	}

	_, err := gateway.Template.GenerateEndpointRegisterFile(
		gateway.EndpointModules, gateway.PackageHelper,
	)
	return err
}

// GenerateMain will generate the main files for the gateway
func (gateway *GatewaySpec) GenerateMain() error {
	_, err := gateway.Template.GenerateMainFile(
		gateway, gateway.PackageHelper,
	)
	return err
}

// GenerateMiddlewareSchemas will generate the middleware schema files
// for the gateway
func (gateway *GatewaySpec) GenerateMiddlewareSchemas() error {

	// Validate middleware go files.

	//for _, middlewares := range gateway.MiddlewareModules {
	//
	//}

	return nil
}
