BINARY=terraform-provider-anchore
TEST?=$$(go list ./...)
GOFMT_FILES?=$$(find . -name '*.go')

default: build

build: fmtcheck
	go install

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
