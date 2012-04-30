#!/bin/bash

go clean
go build ./...
go build .

port=$1

if [ "z" == "z$1" ]; then
	port=:"8888"
fi

if [ -x "./getdone" ]; then
	./getdone -port $port -root ./web/site/
fi
