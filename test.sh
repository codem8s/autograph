#!/usr/bin/env bash

set -e

option=$1

if [[ ${option} == "with-coverage" ]]; then
    echo "Running tests with code coverage"
    has_coverage=true
fi

for d in $(go list ./... | grep -v vendor); do
    if [ "$has_coverage" = "true" ]; then
        echo "Running tests with coverage for $d"
        go test -race -coverprofile=profile.out -covermode=atomic ${d}
    else
        echo "Running tests for $d"
        go test -timeout 3m ${d}
    fi
    if [ -f profile.out ]; then
        cat profile.out >> coverage.txt
        rm profile.out
    fi
done

if [ "$has_coverage" = "true" ]; then
    bash <(curl -s https://codecov.io/bash)
fi