#!/bin/bash

set -e

cd src

echo Building server binary...
go build -o ../out/ server.go

echo Building client binaries...
go build -o ../out/ client-http.go
go build -o ../out/ client-grpc.go

cd ..

echo Done.
echo
echo Binaries are located in "out/" directory:
ls -al out
