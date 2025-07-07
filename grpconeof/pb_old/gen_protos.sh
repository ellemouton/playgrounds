#!/bin/bash

set -e

# Directory of this script regardless of the current working directory.
DIR="$(cd "$(dirname "${BASH_SOURCE[0]}")" && pwd)"

# generate compiles the *.pb.go stubs from the *.proto files.
generate() {
  for file in "${DIR}"/*.proto; do
    protoc -I/usr/local/include -I"${DIR}" \
      --go_out=plugins=grpc,paths=source_relative:"${DIR}" \
      "${file}"
  done
}

# format formats the *.proto files with clang-format.
format() {
  find "${DIR}" -name "*.proto" -print0 | \
    xargs -0 clang-format --style=file -i
}

# Execute the steps from the script's directory so it works no matter
# where it's called from.
cd "${DIR}"
format
generate
