#!/bin/sh
set -o errexit
# TARGET='vendor/github.com/peartherapeutics/bitrise-step-api-build-trigger'
set -x
# Monkey-patch swagger definition to include 201 response to work around strict checks in generated client
curl 'https://api-docs.bitrise.io/docs/swagger.json' | jq 'del(.paths."/apps/{app-slug}/builds".post.responses."200") | .paths."/apps/{app-slug}/builds".post.responses."201" |= {"description":"Created","schema":{"type":"object","$ref":"#/definitions/v0.BuildTriggerRespModel"}}' > swagger.json
rm -Rf models operations client
swagger generate client #--target="$TARGET"
