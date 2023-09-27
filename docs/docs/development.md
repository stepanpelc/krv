---
layout: default
title: Development
nav_order: 4
---

# Development

---
{: .no_toc }

<details open markdown="block">
  <summary>
    Table of contents
  </summary>
  {: .text-delta }
- TOC
{:toc}
</details>

---

## Runtime

- container port exposed `8080`
- kubernetes service cluster port `5000` e.g. krv.krv-system:5000
- health check path `/health`
- rest path for actual Validation statues `/validations`

## Build

### Native build

Prerequisities:

- at least GO v16 must be installed

```go
CGO_ENABLED=0 go build
```

### Docker build

```bash
docker build -t krv:latest . -f ./Dockerfile
```

## Install

### Quick install 

`helm install krv  --namespace krv-system --values values.yaml .`  

### Customized install
  
`helm install krv  --namespace krv-system --values values.yaml . --set timeInterval=<time_in_minutes> --set logLevel=<error,warn,info,debug,trace>`

## Project structure

- `src/api/crd/v1` - Validation CRD v1 objects, structures and openapi definition
- `src/client` - definition and initializtion of k8s api-server clients. Besides standard k8s clients (kubernetes, apiextensions, dynamic) it define and register our custom CRD client
- `src/server/http` http listener handle requests. TEST api for health check and Validations status getter is defined here
- `src/shared` - common variables and constans shared across another packages
- `src/watcher` - main logic of krv. Periodically run validations of defined resources. Also use  watch-cache which is used for http GET operations, so it is not necessary to hit api-server everytime new incoming request comes
- `helmcharts` - helm manifests
- `docs` - documentation
