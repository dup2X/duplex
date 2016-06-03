#!/bin/bash

rm -rf coverage.out
go test -coverprofile=coverage.out
#go tool cover -func=coverage.out
go tool cover -html=coverage.out
