.PHONY: test

test:
	@go test -v ./pkg/...

fmt:
	@gofmt -w *.go

vet:
	@go vet




