#!/bin/bash
CWD="$( cd "$( dirname "${BASH_SOURCE[0]}" )" && pwd )"
export GOBIN=${CWD}/bin
export GOPATH=`godep path`:$GOPATH

#go install github.com/rkrombho/getchctl

cd $CWD/getchctl
GOOS=linux GOARCH=386 go install
GOOS=linux GOARCH=amd64 go install
GOOS=windows GOARCH=386 go install
GOOS=windows GOARCH=amd64 go install


