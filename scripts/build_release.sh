#!/usr/bin/env bash

set -euo pipefail

#cd "${0%/*}"

version="`git tag | tail -1`"
go build -ldflags "-X main.buildVersion=${version} -s -w"

