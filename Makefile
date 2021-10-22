REGISTRY=registry.terraform.io
NAMESPACE=excellaco
NAME=anchore
VERSION=$(shell git describe --tags --abbrev=0 | sed 's/^v//')
BINARY=terraform-provider-$(NAME)_$(VERSION)
INSTALL_PATH=${HOME}/.terraform.d/plugins/$(REGISTRY)/$(NAMESPACE)/$(NAME)/$(VERSION)
TEST?=$$(go list ./...)
GOFMT_FILES?=$$(find . -name '*.go')

default: install

build: fmtcheck
	GOOS=darwin GOARCH=amd64 go build -o $(BINARY)_darwin_amd64
	GOOS=linux GOARCH=amd64 go build -o $(BINARY)_linux_amd64

install: build
	mkdir -p $(INSTALL_PATH)
	mkdir -p $(INSTALL_PATH)/darwin_amd64 && mv $(BINARY)_darwin_amd64 $(INSTALL_PATH)/darwin_amd64/$(BINARY)
	mkdir -p $(INSTALL_PATH)/linux_amd64 && mv $(BINARY)_linux_amd64 $(INSTALL_PATH)/linux_amd64/$(BINARY)

test:
	go test $(TEST) || exit 1
	echo $(TEST) | \
		xargs -t -n4 go test $(TESTARGS) -timeout=30s -parallel=4

clean:
	rm -f "$(shell go env GOPATH)/bin/$(BINARY)"

fmt:
	gofmt -s -w $(GOFMT_FILES)

fmtcheck:
	@sh -c "'$(CURDIR)/scripts/gofmtcheck.sh'"
