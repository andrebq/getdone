#!/bin/bash

go clean
go build ./...
go build .
./getdone -port $1 -root ./web/site/
