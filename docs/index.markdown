---
# Feel free to add content and custom Front Matter to this file.
# To modify the layout, see https://jekyllrb.com/docs/themes/#overriding-theme-defaults

title: krv - Kubernetes Resource Validator
author: Tao He
layout: home
---
Tool designed to simplify the validation of Kubernetes resources using the power of JSONPath. The `krv` provides a user-friendly and efficient solution to validate your in-cluster resources against custom-defined criteria.

With `krv`, you can harness the flexibility of JSONPath expressions to specify complex validation rules, making it easy to enforce best practices, security policies, and compliance standards across your Kubernetes clusters. Say goodbye to manual resource inspection and hello to automated, error-free validation with `krv`.

Key Features:

- Intuitive JSONPath-based validation: Create rules using JSONPath expressions to pinpoint specific fields in your Kubernetes resources.
- Easy integration: Seamlessly integrate `krv` into your CI/CD pipelines or Kubernetes workflows for automated resource validation.
- Comprehensive error reporting: Receive clear and actionable feedback on validation results, helping you quickly identify and resolve issues.
- Extensive resource support: Validate any Kubernetes resource.

The `krv` empowers you to ensure the quality and compliance of your Kubernetes resources effortlessly, saving you time and reducing the risk of misconfigurations.

Installation

```bash
helm install <my-release> oci://registry-1.docker.io/stepanpelc/krv-helm
```

Usage

```yaml
apiVersion: krv.sizek.cz/v1
kind: Validation
metadata:
  name: example-validation-one
spec:
  name: ""
  namespace: "kube-system"
  resource: "Configmap"
  validation:
    - jsonPath: "data.cluster_env"
      value: "NPROD"
    - jsonPath: "data.tenant_prefix"
      value: "ps-sq002"
```