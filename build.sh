#!/bin/bash
DIR=$( cd "$( dirname "${BASH_SOURCE[0]}" )" && pwd )
export GOPATH=$DIR
go install cli
cp -v bin/cli /usr/bin/wucli