export CGO_ENABLED ?= 0
GOFLAGS += -trimpath
LDFLAGS += -X main.version=$(VERSION)

builddir = bin
distdir = dist
tmpdir = tmp

all: test build

benchmark: build
	go test ./... -bench . -benchmem -benchtime 10000x

build:
	@mkdir -p "$(builddir)"
	go build $(GOFLAGS) -ldflags "$(LDFLAGS)" -o "$(builddir)/" ./...

check:
	golangci-lint run ./...

clean:
	rm -rf "$(builddir)" "$(distdir)" "$(tmpdir)"

test:
	@mkdir -p "$(tmpdir)/reports"
	go test $(GOFLAGS) -ldflags "$(LDFLAGS)" -coverprofile "$(tmpdir)/reports/coverage.out" ./...
	go tool cover -html "$(tmpdir)/reports/coverage.out" -o "$(tmpdir)/reports/coverage.html"

.PHONY: all build check clean test
