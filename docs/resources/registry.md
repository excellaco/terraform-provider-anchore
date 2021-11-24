---
page_title: "Anchore: anchore_registry"
description: |-
  Provides an Anchore registry.
---

# Resource: anchore_registry

Provides an Anchore registry.

## Example Usage

```terraform
resource "anchore_registry" "ecr" {
  url      = var.ecr_host
  name     = "ECR"
  username = "awsauto"
  password = "awsauto"
  type     = "awsecr"
}
```

## Argument Reference

The following arguments are supported:

* `name` - (Required) A user-supplied name for the registry.
* `url` - (Required) Endpoing of the target container registry.
* `username` - (Required) Login name for the container registry.
* `password` - (Required) Login password for the container registry.
* `type` - (Optional) Type of container registry.
