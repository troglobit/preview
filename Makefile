VERSION ?= $(shell git describe --tags --always --dirty 2>/dev/null || echo dev)
LDFLAGS  = -s -w -X github.com/troglobit/preview/cmd.version=$(VERSION)

all: build

build:
	go build -ldflags "$(LDFLAGS)" -o preview .

run: build
	./preview $(ARGS)

test:
	go test ./...

vet:
	go vet ./...

fmt:
	gofmt -w .

fmt-check:
	@test -z "$$(gofmt -l .)" || { gofmt -l .; echo "please run make fmt"; exit 1; }

lint:
	golangci-lint run

clean:
	rm -f preview

.PHONY: all build run test vet fmt fmt-check lint clean
