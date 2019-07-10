#!/bin/sh
set -o errexit
# TARGET='vendor/github.com/peartherapeutics/bitrise-step-api-build-trigger'
set -x
curl -O 'https://api-docs.bitrise.io/docs/swagger.json'
# mkdir -p "$TARGET"
swagger generate client #--target="$TARGET"
