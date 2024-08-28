#!/bin/bash

PRJ_DIR="$(cd "$(dirname "${BASH_SOURCE[0]}")" && pwd)"
pushd ${PRJ_DIR}

PATH=/mnt/c/bin/protoc-26.1-linux-x86_64/bin:${HOME}/go/bin:${PATH}
APP_COMMIT=$(git rev-parse --short HEAD)
APP_BUILD_TIME=$(date -u '+%Y-%m-%d_%H:%M:%S')

go mod tidy
gofmt -w .
go vet http_get_files
go test http_get_files
go run http_get_files
