# Go API client for ionoscloud

Logging Service is a service that provides a centralized logging system where users are able to push and aggregate their system or application logs. This service also provides a visualization platform where users are able to observe, search and filter the logs and also create dashboards and alerts for their data points.
This service can be managed through a browser-based \"Data Center Designer\" (DCD) tool or via an API.
The API allows you to create logging pipelines or modify existing ones. It is designed to allow users to leverage the same power and flexibility found within the DCD visual tool. Both tools are consistent with their concepts and lend well to making the experience smooth and intuitive.

## Overview
This API client was generated by the [OpenAPI Generator](https://openapi-generator.tech) project.  By using the [OpenAPI-spec](https://www.openapis.org/) from a remote server, you can easily generate an API client.

- API version: 0.0.1
- Package version: 1.0.0
- Build package: org.openapitools.codegen.languages.GoClientCodegen

## Installation

Install the following dependencies:

```shell
go get github.com/stretchr/testify/assert
go get golang.org/x/oauth2
go get golang.org/x/net/context
```

Put the package under your project folder and add the following in import:

```golang
import sw "./ionoscloud"
```

To use a proxy, set the environment variable `HTTP_PROXY`:

```golang
os.Setenv("HTTP_PROXY", "http://proxy_name:proxy_port")
```

## Configuration of Server URL

Default configuration comes with `Servers` field that contains server objects as defined in the OpenAPI specification.

### Select Server Configuration

For using other server than the one defined on index 0 set context value `sw.ContextServerIndex` of type `int`.

```golang
ctx := context.WithValue(context.Background(), sw.ContextServerIndex, 1)
```

### Templated Server URL

Templated server URL is formatted using default variables from configuration or from context value `sw.ContextServerVariables` of type `map[string]string`.

```golang
ctx := context.WithValue(context.Background(), sw.ContextServerVariables, map[string]string{
	"basePath": "v2",
})
```

Note, enum values are always validated and all unused variables are silently ignored.

### URLs Configuration per Operation

Each operation can use different server URL defined using `OperationServers` map in the `Configuration`.
An operation is uniquely identifield by `"{classname}Service.{nickname}"` string.
Similar rules for overriding default operation server index and variables applies by using `sw.ContextOperationServerIndices` and `sw.ContextOperationServerVariables` context maps.

```
ctx := context.WithValue(context.Background(), sw.ContextOperationServerIndices, map[string]int{
	"{classname}Service.{nickname}": 2,
})
ctx = context.WithValue(context.Background(), sw.ContextOperationServerVariables, map[string]map[string]string{
	"{classname}Service.{nickname}": {
		"port": "8443",
	},
})
```

## Documentation for API Endpoints

All URIs are relative to *https://logging.de-txl.ionos.com*

Class | Method | HTTP request | Description
------------ | ------------- | ------------- | -------------
*PipelinesApi* | [**PipelineKey**](docs/api/PipelinesApi.md#pipelinekey) | **Post** /pipelines/{pipelineId}/key | Renews the key of a Pipeline
*PipelinesApi* | [**PipelinesDelete**](docs/api/PipelinesApi.md#pipelinesdelete) | **Delete** /pipelines/{pipelineId} | Delete a pipeline
*PipelinesApi* | [**PipelinesFindById**](docs/api/PipelinesApi.md#pipelinesfindbyid) | **Get** /pipelines/{pipelineId} | Fetch a pipeline
*PipelinesApi* | [**PipelinesGet**](docs/api/PipelinesApi.md#pipelinesget) | **Get** /pipelines | List pipelines
*PipelinesApi* | [**PipelinesPatch**](docs/api/PipelinesApi.md#pipelinespatch) | **Patch** /pipelines/{pipelineId} | Patch a pipeline
*PipelinesApi* | [**PipelinesPost**](docs/api/PipelinesApi.md#pipelinespost) | **Post** /pipelines | Create a pipeline


## Documentation For Models

 - [CreateRequest](docs/models/CreateRequest.md)
 - [CreateRequestPipeline](docs/models/CreateRequestPipeline.md)
 - [CreateRequestProperties](docs/models/CreateRequestProperties.md)
 - [Destination](docs/models/Destination.md)
 - [ErrorMessage](docs/models/ErrorMessage.md)
 - [ErrorResponse](docs/models/ErrorResponse.md)
 - [InlineResponse200](docs/models/InlineResponse200.md)
 - [Metadata](docs/models/Metadata.md)
 - [PatchRequest](docs/models/PatchRequest.md)
 - [PatchRequestPipeline](docs/models/PatchRequestPipeline.md)
 - [PatchRequestProperties](docs/models/PatchRequestProperties.md)
 - [Pipeline](docs/models/Pipeline.md)
 - [PipelineListResponse](docs/models/PipelineListResponse.md)
 - [PipelineProperties](docs/models/PipelineProperties.md)
 - [Processor](docs/models/Processor.md)
 - [ResponsePipeline](docs/models/ResponsePipeline.md)
 - [ResponsePipelineAllOf](docs/models/ResponsePipelineAllOf.md)
 - [ResponsePipelineAllOf1](docs/models/ResponsePipelineAllOf1.md)


## Documentation For Authorization



### tokenAuth

- **Type**: API key
- **API key parameter name**: Authorization
- **Location**: HTTP header

Note, each API key must be added to a map of `map[string]APIKey` where the key is: Authorization and passed in as the auth context for each request.


## Documentation for Utility Methods

Due to the fact that model structure members are all pointers, this package contains
a number of utility functions to easily obtain pointers to values of basic types.
Each of these functions takes a value of the given basic type and returns a pointer to it:

* `PtrBool`
* `PtrInt`
* `PtrInt32`
* `PtrInt64`
* `PtrFloat`
* `PtrFloat32`
* `PtrFloat64`
* `PtrString`
* `PtrTime`

## Author


