#!/usr/bin/env bash
# https://github.com/codecov/example-go#caveat-multiple-files

ee() {
  echo "+ $*"
  eval "$@"
}

cd "$(dirname "$0")/.." || exit 1

echo "" > coverage.txt || exit 1

for d in $(go list ./...); do
  echo "$d" || exit 1
  ee go test -race -coverprofile=profile.out -covermode=atomic "$d" || exit 1
  if [ -f profile.out ]; then
    cat profile.out >> coverage.txt || exit 1
    rm profile.out || exit 1
  fi
done
