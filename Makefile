.PHONY: deps lint

# Check for required command tools to build or stop immediately
EXECUTABLES = git go find pwd
K := $(foreach exec,$(EXECUTABLES),\
        $(if $(shell which $(exec)),some string,$(error "No $(exec) in PATH)))

GO ?= latest

ROBOTNAME = qlcrobot
ROBOTOMAIN = $(shell pwd)/example/robot/main.go

BUILDDIR = $(shell pwd)/build
VERSION ?= $(shell git describe --tags `git rev-list --tags --max-count=1`)
GITREV = $(shell git rev-parse --short HEAD)
BUILDTIME = $(shell date +'%FT%TZ%z')
LDFLAGS=-ldflags "-X main.version=${VERSION} -X main.commit=${GITREV} -X main.date=${BUILDTIME}"

deps:
	go get -u github.com/git-chglog/git-chglog/cmd/git-chglog

lint: 
	golangci-lint run --fix
