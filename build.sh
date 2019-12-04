#!/bin/bash
export PATH=/usr/local/go/bin:$PATH
export GOPATH=/root/work/
go build -o /build/go/bin/main /build/go/src/main.go
