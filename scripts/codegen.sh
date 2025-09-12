#!/usr/bin/env bash

set -euo pipefail
set -x

repository_root="$(git rev-parse --show-toplevel)"

cd "$repository_root/cmd/codegen"

go run "$repository_root/cmd/codegen" \
	-package ini \
	-method-receiver ps -method-type '*PropSet' \
	-option-name property -option-article a \
	-out "$repository_root/ini/gen.go"

go run "$repository_root/cmd/codegen" \
	-package env \
	-method-receiver es -method-type '*EnvSet' \
	-option-name "env var" -option-article an \
	-out "$repository_root/env/gen.go"

go run "$repository_root/cmd/codegen" \
	-package configue \
	-method-receiver f -method-type '*Figue' \
	-option-name option -option-article an \
	-out "$repository_root/gen.go"

cd "$repository_root"
go fmt "./..."
