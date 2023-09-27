---
layout: default
title: Validations
nav_order: 2
---

# Validations

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

## Validation custom resource

Structure of validation resource

```yaml
apiVersion: krv.sizek.cz/v1
kind: Validation
metadata:
  name: <validation name>
spec:
  name: "<resource name>"
  namespace: "<resource namespace>"
  resource: "<resource kind>"
  validation:
    - jsonPath: "<jsonPath>"
      value: "<value>"
```

Detailed description:

- `metadata.name` - Name of validation resource
- `spec.name` - Name of validated resource (can be defined as regex)
- `spec.namespace` - Namespace of validated resource
- `spec.resource` - Kind of validated resource
- `spec.validations` - Array of validations
- `spec.validations.jsonPath` - jsonPath of validated attribute
- `spec.validations.value` - value of validated attribute

{: .note }
> You can use jq tool for finding of corect jsonPath in validated resource.
>
> `kubectl get pod -n kube-system etcd-kind-control-plane -ojson | jq '.spec.restartPolicy'`

## Validation outputs

### `kubectl`

Validations are presented via kubectl command-line tool as list of objects

```bash
NAME              RESOURCE-NAMESPACE   RESOURCE   STATE     AGE
etcd-pullpolicy   kube-system          Pod        OK        5m24s
```

Validations are checked in configured interval (default 5 minutes).

### States of validations

- **OK** - attribute(s) with specified value found
- **NOK** - attribute(s) with specified value not found
- **MISSING** - resource not found

## Examples

### Existing resource

```yaml
apiVersion: krv.sizek.cz/v1
kind: Validation
metadata:
  name: etcd-pullpolicy
spec:
  name: "etcd-kind-control-plane"
  resource: "Pod"
  namespace: "kube-system"
  validation: []
```

### Resource with attribute

```yaml
apiVersion: krv.sizek.cz/v1
kind: Validation
metadata:
  name: etcd-pullpolicy
spec:
  name: "etcd-kind-control-plane"
  resource: "Pod"
  namespace: "kube-system"
  validation:
    - jsonPath: "spec.restartPolicy"
      value: "Always"
```

### Resource with multiple attributes

```yaml
apiVersion: krv.sizek.cz/v1
kind: Validation
metadata:
  name: etcd-pullpolicy
spec:
  name: "etcd-kind-control-plane"
  resource: "Pod"
  namespace: "kube-system"
  validation:
    - jsonPath: "spec.restartPolicy"
      value: "Always"
    - jsonPath: "spec.hostNetwork"
      value: "true"

```
