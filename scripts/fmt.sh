#!/usr/bin/env sh

cd "$(dirname "$0")/.." || exit 1

find . \
  -type d -name .git -prune -o \
  -type f -name "*.go" -print0 |
  xargs -0 gofmt -l -s -w
