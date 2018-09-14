#!/usr/bin/env bash

set -euo pipefail

#cd "${0%/*}"

buildInfo="`date -u '+%Y-%m-%dT%TZ'`|`git describe --always --long`|`git tag | tail -1`"
go build -ldflags "-X main.buildInfo=${buildInfo} -s -w"

