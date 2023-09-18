# Development Guide

## Runtime

- container port exposed `8080`
- kubernetes service cluster port `5000` e.g. krv.krv-system:5000
- health check path `/health`
- rest path for actual Validation statues `/validations`

## Build

native build
- at least GO v16 must be installed
```go
CGO_ENABLED=0 go build
```

docker build
```azure
docker build -t krv:latest . -f ./Dockerfile
```

## Install

`helm install krv  --namespace krv-system --values values.yaml .`  
or  
`helm install krv  --namespace krv-system --values values.yaml . --set timeInterval=<time_in_minutes> --set logLevel=<error,warn,info,debug,trace>`

## Project structure

`api/crd/v1` Validation CRD v1 objects, structures and openapi definition

`client` definition and initializtion of k8s api-server clients. Besides standard k8s clients (kubernetes, apiextensions, dynamic) it define and register our custom CRD client

`server/http` http listener handle requests. TEST api for health check and Validations status getter is defined here

`shared` common variables and constans shared across another packages

`watcher` main logic of krv. Periodically run validations of defined resources. Also use  watch-cache which is used 
for http GET operations, so it is not necessary to hit api-server everytime new incoming request comes 

`example` examples of Validation Kubernetes resource definitions

`helmcharts` helm manifests