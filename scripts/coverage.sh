#!/usr/bin/env sh

cd "$(dirname "$0")"/.. || exit 1

mkdir -p ".coverage/$1" || exit 1
go test "./$1" -coverprofile=".coverage/$1/coverage.txt" -covermode=atomic || exit 1
go tool cover -html=".coverage/$1/coverage.txt"
