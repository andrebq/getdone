#!/bin/bash

go clean

go test ./...
testResult=$?
if [ "z0" != "z$testResult" ]; then
	echo "Tests failed can't proceed. Exit code: $testResult"
	exit 1
fi

go build ./...
go build .

port=$1

if [ "z" == "z$1" ]; then
	port=:"8888"
fi

if [ -x "./getdone" ]; then
	./getdone -port $port -root ./web/site/
fi
