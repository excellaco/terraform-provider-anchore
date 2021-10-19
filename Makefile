BINARY=terraform-provider-anchore

default: build

build:
	go install

test:
	go test -cover ./...

clean:
	rm -f "$(shell go env GOPATH)/bin/$(BINARY)"
