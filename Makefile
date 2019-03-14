.PHONY: all clean
.PHONY: robot_linux  robot-linux-amd64 robot-darwin-amd64
.PHONY: robot-darwin robot-darwin-amd64
.PHONY: robot-windows robot-windows-386 robot-windows-amd64


# Check for required command tools to build or stop immediately
EXECUTABLES = git go find pwd
K := $(foreach exec,$(EXECUTABLES),\
        $(if $(shell which $(exec)),some string,$(error "No $(exec) in PATH)))

GO ?= latest

ROBOTNAME = qlcrobot
ROBOTOMAIN = $(shell pwd)/example/robot/main.go

BUILDDIR = $(shell pwd)/build
VERSION = 0.0.1
GITREV = $(shell git rev-parse --short HEAD)
BUILDTIME = $(shell date +'%Y-%m-%d_%T')
LDFLAGS=-ldflags "-X main.version=${VERSION} -X main.sha1ver=${GITREV} -X main.buildTime=${BUILDTIME}"

build:
	@echo "package qlcchain" > $(shell pwd)/version.go
	@echo  "">> $(shell pwd)/version.go
	@echo "const GITREV = \""$(GITREV)"\"" >> $(shell pwd)/version.go
	@echo "const VERSION = \""$(VERSION)"\"" >> $(shell pwd)/version.go
	@echo "const BUILDTIME = \""$(BUILDTIME)"\"" >> $(shell pwd)/version.go
	go build ${LDFLAGS} -v -i -o $(BUILDDIR)/$(ROBOTNAME) $(ROBOTOMAIN)
	@echo "Build robot done."

all: robot-windows robot-darwin robot-linux

clean:
	rm -rf $(BUILDDIR)/

robot-linux: robot-linux-amd64
	@echo "Linux cross compilation done:"

robot-linux-amd64:
	env GOOS=linux GOARCH=amd64 go build ${LDFLAGS} -i -o $(BUILDDIR)/$(ROBOTNAME)-linux-amd64-v$(VERSION)-$(GITREV) $(ROBOTOMAIN)
	@echo "Build linux amd64 done."
	@ls -ld $(BUILDDIR)/$(ROBOTNAME)-linux-amd64-v$(VERSION)-$(GITREV)

robot-darwin:
	env GOOS=darwin GOARCH=amd64 go build ${LDFLAGS} -i -o $(BUILDDIR)/$(ROBOTNAME)-darwin-amd64-v$(VERSION)-$(GITREV) $(ROBOTOMAIN)
	@echo "Build darwin server done."
	@ls -ld $(BUILDDIR)/$(ROBOTNAME)-darwin-amd64-v$(VERSION)-$(GITREV)

robot-windows: robot-windows-amd64 robot-windows-386
	@echo "Windows cross compilation done:"
	@ls -ld $(BUILDDIR)/$(ROBOTNAME)-windows-*

robot-windows-386:
	env GOOS=windows GOARCH=386 go build ${LDFLAGS} -i -o $(BUILDDIR)/$(ROBOTNAME)-windows-386-v$(VERSION)-$(GITREV).exe $(ROBOTOMAIN)
	@echo "Build windows x86 done."
	@ls -ld $(BUILDDIR)/$(ROBOTNAME)-windows-386-v$(VERSION)-$(GITREV).exe

robot-windows-amd64:
	env GOOS=windows GOARCH=amd64 go build ${LDFLAGS} -i -o $(BUILDDIR)/$(ROBOTNAME)-windows-amd64-v$(VERSION)-$(GITREV).exe $(ROBOTOMAIN)
	@echo "Build windows amd64 done."
	@ls -ld $(BUILDDIR)/$(ROBOTNAME)-windows-amd64-v$(VERSION)-$(GITREV).exe

