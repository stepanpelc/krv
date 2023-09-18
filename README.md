# Kubernetes Resource Validator

## Intro

Kubernetes Resource Validator hereinafter `krv` is tool for checking state of kubernetes resources based on `validation` CRD.

## Details

- all checks are created as CRD manifests - `validations` \ `validation`
- internal state of all components is writen into `validation` custom resource
- `validations` can be listed via `kubectl` with resource api shortcode (`sc` instead of `storage-class`)
- `validations` are deployed into same namespace as `krv` application
- `validations` must be easily added and removed

### Example of `kubectl` output

```bash
kubectl get validations --all-namespaces
NAME          RESOURCE NAMESPACE  RESOURCE    NAME            STATE     AGE
pod-check     kube-system         pod         core-dns*.      MISSING   35d
sc-check                          sc          nfs-ganesha-sc  OK         4d
nfs-sc        nfs-test            pvc         nfs-test        NOK        5h
test-deploy   test-app            deployment  net-checker     OK         6h
```

- validations can be applied to **any** kubernetes resources
- validation is free-form definition with posibility of regexp

### Example 1 of validation definition part

```yaml
...
    resource: persistent-storage-class
    name: nfs-test
    namespace: nfs-test
    validation:
    - jsonPath: "status.phase"
      value: "Bound"
...
```

### Example 2 of validation definition part

```yaml
...
    resource: deployment
    name: "test-deploy-[0-9]*"
    namespace: test-app
    validation:
    - jsonPath: "spec.replicas"
      value: "3"
    - jsonPath: "status.availableReplicas"
      value: "3"
...
```

Application itself provides read-only API for state of `validations` in the form of json payload (same as `kubectl get validations -ojson`) for REST only access.

## Happy day scenario

### Start

- `validation` CRD is deployed into cluster
- `krv` is deployed into `krv-system` namespace
- `krv` starts and checks existence of `validation` CRD
- `krv` loads `validations` from `krv-system` namespace
- `krv` watches `validations` resources

### Checks

- in case of existing validated resource check is performed against validation part of definition and set to OK / NOK state
- in case of non-existing resource validation state is set to MISSING state
- in all cases `last check` / `check change` time is updated
- in case of state change additional `event` is writen to `krv-system` namespace

## Parts of delivery

- `validation` CRD
- `krv` application
- RBAC model for `krv` application
- service for `krv` API
- Helm Chart package

## Examples from `kubernetes-client` library

- [watch for resources](https://github.com/kubernetes-client/python/blob/master/examples/watch/pod_namespace_watch.py)
- [usage of `CRD` resources](https://github.com/kubernetes-client/python/blob/master/examples/namespaced_custom_object.py)
