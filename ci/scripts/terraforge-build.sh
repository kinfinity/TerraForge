#!/bin/bash

set -eu

USAGE="USAGE:
${0}"

if [[ $# -ne 0 ]]; then
    echo "${USAGE}" >&2
    exit 1
fi

# BUILD PLATFORM ARCHITECTURE
LINUX_ARCH=amd64
LINUX_OS=linux

# - script executes within its directory
# - prevent breaking in different locations
SCRIPT_DIR="$(dirname "${0}")"
pushd ${SCRIPT_DIR} > /dev/null


# 
TERRAFORGE_MAIN_DIR="cmd"
RUN_DIR="$(cd "../../${TERRAFORGE_MAIN_DIR}"; pwd -P)"
pushd ${RUN_DIR} > /dev/null

# BUILD TERRAFORGE
GOOS=$LINUX_OS GOARCH=$LINUX_ARCH go build -ldflags="-s -w" -o ../build/terraforge terraforge.go