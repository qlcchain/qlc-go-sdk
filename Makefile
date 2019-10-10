.PHONY: deps clean build lint snapshot changelog release

# Check for required command tools to build or stop immediately
EXECUTABLES = git go find pwd
K := $(foreach exec,$(EXECUTABLES),\
        $(if $(shell which $(exec)),some string,$(error "No $(exec) in PATH)))

GO ?= latest

ROBOTNAME = qlcrobot
ROBOTOMAIN = $(shell pwd)/example/robot/main.go

BUILDDIR = $(shell pwd)/build
VERSION ?= 1.2.4
GITREV = $(shell git rev-parse --short HEAD)
BUILDTIME = $(shell date +'%FT%TZ%z')
LDFLAGS=-ldflags "-X main.version=${VERSION} -X main.commit=${GITREV} -X main.date=${BUILDTIME}"

deps:
	go get -u github.com/golangci/golangci-lint/cmd/golangci-lint
	go get -u github.com/git-chglog/git-chglog/cmd/git-chglog
	go get -u github.com/goreleaser/goreleaser

build:
	@echo "package qlcchain" > $(shell pwd)/version.go
	@echo  "">> $(shell pwd)/version.go
	@echo "const GITREV = \""$(GITREV)"\"" >> $(shell pwd)/version.go
	@echo "const VERSION = \""$(VERSION)"\"" >> $(shell pwd)/version.go
	@echo "const BUILDTIME = \""$(BUILDTIME)"\"" >> $(shell pwd)/version.go
	go build ${LDFLAGS} -v -i -o $(BUILDDIR)/$(ROBOTNAME) $(ROBOTOMAIN)
	@echo "Build robot done."

clean:
	rm -rf $(BUILDDIR)/

lint: 
	golangci-lint run --fix

snapshot:
	goreleaser --snapshot --rm-dist

changelog:
	git-chglog $(VERSION) > CHANGELOG.md

release: changelog
	goreleaser --rm-dist --skip-publish --release-notes=CHANGELOG.md