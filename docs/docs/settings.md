---
layout: default
title: Settings
nav_order: 3
---

# Settings

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

## Helm Chart values

Structure of validation resource

```yaml
# Default values for krv.
# This is a YAML-formatted file.
# Declare variables to be passed into your templates.

replicaCount: 1

image:
  repository:  sizek/krv
  tag: 0.1.1
  pullPolicy: Always

imagePullSecrets: []
nameOverride: ""
fullnameOverride: ""

serviceAccount:
  # The name of the service account to use.
  # If not set and create is true, a name is generated using the fullname template
  name: "krv"

podSecurityContext: {}
  # fsGroup: 2000

securityContext:
  # capabilities:
  #   drop:
  #   - ALL
  # readOnlyRootFilesystem: true
  # runAsNonRoot: true
  runAsUser: 1000

service:
  type: ClusterIP
  port: 5000

resources: {}
  # We usually recommend not to specify default resources and to leave this as a conscious
  # choice for the user. This also increases chances charts run on environments with little
  # resources, such as Minikube. If you do want to specify resources, uncomment the following
  # lines, adjust them as necessary, and remove the curly braces after 'resources:'.
  # limits:
  #   cpu: 100m
  #   memory: 128Mi
  # requests:
  #   cpu: 100m
  #   memory: 128Mi

nodeSelector: {}

tolerations: []

affinity: {}

logLevel: "info"

#time in minutes.
timeInterval: 5
```

## Attributes

- `service.type` - krv reporting kubernetes service type
- `service.port` - krv reporting kubernetes service port
- `logLevel` - krv log level
- `timeInterval` - checking interval