#!/bin/sh

go test -v -coverpkg=./internal/... -coverprofile=./profile.cov ./...
go tool cover -func ./profile.cov | grep total | awk '{print $3}'
test_result=$?
exit $test_result
