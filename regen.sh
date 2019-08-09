#!/bin/bash

# This script regenerates go code for google/api/annotations.proto

set -e

GOOGLEAPIS_REPO=https://github.com/googleapis/googleapis
remove_dirs=
trap 'rm -rf $remove_dirs' EXIT

if [ -z "$GOOGLEAPIS" ]; then
  apidir=$(mktemp -d -t regen-cds-api.XXXXXX)
  git clone $GOOGLEAPIS_REPO $apidir
  remove_dirs="$remove_dirs $apidir"
else
  apidir="$GOOGLEAPIS"
fi

# Install protoc-gen-go V2.
go get -u google.golang.org/protobuf/cmd/protoc-gen-go

# Invoke protoc to generate go code.
protoc --plugin=protoc-gen-go=$HOME/go/bin/protoc-gen-go --go_out=. -I $apidir $apidir/google/api/*.proto

# Sanity check the build.
echo 1>&2 "Checking that the libraries build..."
go build -v ./...

gofmt -s -l -w . && goimports -w .

echo 1>&2 "All done!"
