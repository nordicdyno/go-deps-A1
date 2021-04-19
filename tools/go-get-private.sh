#!/usr/bin/env bash
set -e
echo "patch global git config"
git config --global url."git@github.com:nordicdyno/go-deps-private".insteadOf "https://github.com/nordicdyno/go-deps-private"

export GOPRIVATE=github.com/nordicdyno/go-deps-private
set -x
go get -v github.com/nordicdyno/go-deps-private
go mod tidy
set +x

echo "rollback patch of global git config"
git config --global --unset url."git@github.com:nordicdyno/go-deps-private".insteadOf "https://github.com/nordicdyno/go-deps-private"
