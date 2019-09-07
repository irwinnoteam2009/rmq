export GO111MODULE=on
export GOFLAGS=-mod=vendor

GO ?= go
GO_PACKAGE ?= $$(go list -m)
GIT_REV := $$(git rev-parse --short HEAD)
GIT_VER := $$(git describe --abbrev=0 --tags)
OUT := bin/$(GO_PACKAGE)

default: run

prepare: prepare.lint_tools

prepare.lint_tools:
	@$(GO) get -u github.com/golangci/golangci-lint/cmd/golangci-lint

clean:
	rm -rf ./bin/*

lint:
	golangci-lint run

test:
	@$(GO) test -race -timeout 2m ./...

build: clean
	@$(GO) build -ldflags "-X main.Revision=$(GIT_REV) -X main.Version=$(GIT_VER)" -o $(OUT) ./internal/cmd

build.win: OUT = bin/rmq.exe
build.win: build 

run: build
	$(OUT)

