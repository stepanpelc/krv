---
title: Development
layout: post
date: 2023-08-25

---

## Prerequisities

- golang v16 or later

## Build

native build

```bash
CGO_ENABLED=0 go build
```

docker build

```bash
docker build -t krv:latest . -f ./Dockerfile
```

## Project structure

- `api/crd/v1` Validation CRD v1 objects, structures and openapi definition

- `client` definition and initializtion of k8s api-server clients. Besides standard k8s clients (kubernetes, apiextensions, dynamic) it define and register our custom CRD client

- `server/http` http listener handle requests. TEST api for health check and Validations status getter is defined here

- `shared` common variables and constans shared across another packages

- `watcher` main logic of krv. Periodically run validations of defined resources. Also use  watch-cache which is used 
for http GET operations, so it is not necessary to hit api-server everytime new incoming request comes 

- `docs` documentation for usage and development

- `helmcharts` helm manifests
