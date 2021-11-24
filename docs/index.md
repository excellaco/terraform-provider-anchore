---
page_title: "Anchore Provider"
subcategory: "Security & Authentication"
description: |-
  Configure the Anchore container scanning engine.  
---

# Anchore Provider

Configure the Anchore container scanning engine.

## Example Usage

```terraform
terraform {
  required_providers {
    anchore = {
      source  = "excellaco/anchore"
      version = "0.1.2"
    }
  }
}

# Configure Anchore
provider "anchore" {
  api_url = "https://anchore.example.com"
}
```

## Argument Reference

In addition to [generic `provider` arguments](https://www.terraform.io/docs/configuration/providers.html)
(e.g., `alias` and `version`), the following arguments are supported in the Anchore
 `provider` block:

* `api_url` - (Optional) This is the Anchore API endpoint. It must be provided, but
  it can also be sourced from the `ANCHORE_CLI_URL` environment variable.

* `username` - (Optional) This is the login user for the Anchore API. It must be provided, but
  it can also be sourced from the `ANCHORE_CLI_USER` environment variable.

* `password` - (Optional) This is the login password for the Anchore API. It must be provided, but
  it can also be sourced from the `ANCHORE_CLI_PASS` environment variable.
