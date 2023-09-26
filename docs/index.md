---
layout: default
title: Home
nav_order: 1
description: "Just the Docs is a responsive Jekyll theme with built-in search that is easily customizable and hosted on GitHub Pages."
permalink: /
---

# krv - Kubernetes resources validator
{: .fs-9 }

Tool designed to simplify the validation of Kubernetes resources using the power of JSONPath and regexp. 
{: .fs-6 .fw-300 }

[Get started now](#quick-start){: .btn .btn-primary .fs-5 .mb-4 .mb-md-0 .mr-2 }
[View it on GitHub][Just the Docs repo]{: .btn .fs-5 .mb-4 .mb-md-0 }

---

 The `krv` provides a user-friendly and efficient solution to validate your in-cluster resources against custom-defined criteria to enforce best practices, security policies, and compliance standards across your Kubernetes clusters. 

Say goodbye to manual resource inspection and hello to automated, error-free validation with `krv`.

Key Features:

- Intuitive JSONPath-based validation: Create rules using JSONPath expressions to pinpoint specific fields in your Kubernetes resources.
- Easy integration: Seamlessly integrate `krv` into your CI/CD pipelines or Kubernetes workflows for automated resource validation.
- Comprehensive error reporting: Receive clear and actionable feedback on validation results, helping you quickly identify and resolve issues.
- Extensive resource support: Validate any Kubernetes resource.

The `krv` empowers you to ensure the quality and compliance of your Kubernetes resources effortlessly, saving you time and reducing the risk of misconfigurations.


## Quick start

Install latest version via `helm`

```bash
helm install <my-release> oci://registry-1.docker.io/stepanpelc/krv-helm
```

Check running pod.

```bash
kubectl get pod
NAMESPACE            NAME                                                READY   STATUS    RESTARTS        AGE
default              <my-release>-krv-helm-6b5fd4bdc7-g25ql              1/1     Running   0               5h39m
```

Check for value in `pod` resource.

```bash
kubectl get pod -n kube-system          etcd-kind-control-plane  -ojson | jq '.spec.containers[0].imagePullPolicy'
"IfNotPresent"
```

Create first validation

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
    - jsonPath: "spec.containers[0].imagePullPolicy"
      value: "IfNotPresent"
```

Apply validation and check existence

```bash
NAME              RESOURCE-NAMESPACE   RESOURCE   STATE     AGE
etcd-pullpolicy   kube-system          Pod                  4s
```

Validations are checked in configured interval (default 5 minutes).


```bash
NAME              RESOURCE-NAMESPACE   RESOURCE   STATE     AGE
etcd-pullpolicy   kube-system          Pod        OK        5m24s
```


## About the project

`krv` is &copy; 2022-{{ "now" | date: "%Y" }} by [SIZEK s.r.o.](https://sizek.cz).

### License

Just the Docs is distributed by an [GPLv3 license](https://github.com/just-the-docs/just-the-docs/tree/main/LICENSE.txt).

### Contributing

When contributing to this repository, please first discuss the change you wish to make via issue,
email, or any other method with the owners of this repository before making a change. Read more about becoming a contributor in [our GitHub repo](https://github.com/sizekcz/krv#contributing).

#### Thank you to the contributors of Just the Docs!

<ul class="list-style-none">
{% for contributor in site.github.contributors %}
  <li class="d-inline-block mr-1">
     <a href="{{ contributor.html_url }}"><img src="{{ contributor.avatar_url }}" width="32" height="32" alt="{{ contributor.login }}"></a>
  </li>
{% endfor %}
</ul>

### Code of Conduct

Just the Docs is committed to fostering a welcoming community.

[View our Code of Conduct](https://github.com/just-the-docs/just-the-docs/tree/main/CODE_OF_CONDUCT.md) on our GitHub repository.

----

[^1]: The [source file for this page] uses all three markup languages.

[^2]: [It can take up to 10 minutes for changes to your site to publish after you push the changes to GitHub](https://docs.github.com/en/pages/setting-up-a-github-pages-site-with-jekyll/creating-a-github-pages-site-with-jekyll#creating-your-site).

[Jekyll]: https://jekyllrb.com
[Markdown]: https://daringfireball.net/projects/markdown/
[Liquid]: https://github.com/Shopify/liquid/wiki
[Front matter]: https://jekyllrb.com/docs/front-matter/
[Jekyll configuration]: https://jekyllrb.com/docs/configuration/
[source file for this page]: https://github.com/just-the-docs/just-the-docs/blob/main/index.md
[Just the Docs Template]: https://just-the-docs.github.io/just-the-docs-template/
[Just the Docs]: https://just-the-docs.com
[Just the Docs repo]: https://github.com/just-the-docs/just-the-docs
[Just the Docs README]: https://github.com/just-the-docs/just-the-docs/blob/main/README.md
[GitHub Pages]: https://pages.github.com/
[Template README]: https://github.com/just-the-docs/just-the-docs-template/blob/main/README.md
[GitHub Pages / Actions workflow]: https://github.blog/changelog/2022-07-27-github-pages-custom-github-actions-workflows-beta/
[customize]: 
[use the template]: https://github.com/just-the-docs/just-the-docs-template/generate
