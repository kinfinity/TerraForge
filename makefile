# Make targets:
#
#  build    : build binary in development mode
#  release  : build binary for RELEASE
#  clean  	: removes all build artifacts
#  test   	: runs tests
VERSION=0.0.1
DOCKER_IMAGE ?= teleport
GIT_TAG=v$(VERSION)


# standard autotools variables
ifneq ("$(wildcard /bin/bash)","")
SHELL := /bin/bash -o pipefail
endif
BUILDDIR ?= build
BINDIR ?= /usr/local/bin
DATADIR ?= /usr/local/share/terraforge
ADDFLAGS ?=
TERRAFORGE_DEBUG ?= false

# 
.PHONY: clean
clean: |
	@echo "---> Cleaning up executables."
	rm -rf $(BUILDDIR)


.PHONY: build
build: 
	@echo "---> Building terraforge executables."
	ci/scripts/terraforge-build.sh $(BUILDDIR) 
