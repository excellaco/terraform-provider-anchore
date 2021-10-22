Terraform Provider for Anchore
==============================

- Website: https://anchore.com/

Building The Provider
---------------------

Clone repository to: `$GOPATH/src/github.com/excellaco/terraform-provider-anchore`
```sh
$ mkdir -p $GOPATH/src/github.com/excellaco
$ cd $GOPATH/src/github.com/excellaco
$ git clone git@github.com:excellaco/terraform-provider-anchore
```

Make sure GO can pull a private Git repository
```sh
$ git config --global url."git@github.com:".insteadOf "https://github.com/"
```

Install the provider
```sh
$ cd $GOPATH/src/github.com/excellaco/terraform-provider-anchore
$ make install
```
