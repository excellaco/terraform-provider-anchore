Terraform Provider for Anchore
==============================

- Website: https://anchore.com/

Building The Provider
---------------------

Clone repository to: `$GOPATH/src/github.com/hashicorp/terraform-provider-anchore`

```sh
$ mkdir -p $GOPATH/src/github.com/excellaco; cd $GOPATH/src/github.com/excellaco
$ git clone git@github.com:excellaco/terraform-provider-anchore
```

Enter the provider directory and build the provider

```sh
$ cd $GOPATH/src/github.com/excellaco/terraform-provider-anchore
$ make build
```
