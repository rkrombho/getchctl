#!/bin/bash
CWD="$( cd "$( dirname "${BASH_SOURCE[0]}" )" && pwd )"
export GOBIN=${CWD}/bin
export GOPATH=`godep path`:$GOPATH

go install github.com/rkrombho/getchctl

