VERSION ?= $(shell git describe --tags --always --dirty 2>/dev/null || echo dev)
LDFLAGS  = -s -w -X github.com/troglobit/preview/src.version=$(VERSION)
PREFIX  ?= /usr/local
MANDIR  ?= $(PREFIX)/share/man

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

lint: man-lint
	golangci-lint run

man-lint:
	mandoc -T lint -W warning doc/preview.1

install: build
	install -D -m 0755 preview $(DESTDIR)$(PREFIX)/bin/preview
	install -D -m 0644 doc/preview.1 $(DESTDIR)$(MANDIR)/man1/preview.1

uninstall:
	rm -f $(DESTDIR)$(PREFIX)/bin/preview $(DESTDIR)$(MANDIR)/man1/preview.1

clean:
	rm -f preview

# Regenerate the theme palettes from @primer/primitives (needs network).
PRIMER_VERSION = 11.10.0
PRIMER = https://unpkg.com/@primer/primitives@$(PRIMER_VERSION)/dist/css/functional/themes
themes:
	@tmp=$$(mktemp -d) && \
	for t in light dark dark-dimmed light-colorblind dark-colorblind; do \
		curl -sfL -o "$$tmp/$$t.css" $(PRIMER)/$$t.css || exit 1; \
	done && \
	python3 -I src/assets/gen-themes.py "$$tmp" $(PRIMER_VERSION) > src/assets/static/css/themes.css && \
	rm -r "$$tmp"

.PHONY: all build run test vet fmt fmt-check lint man-lint install uninstall clean themes
