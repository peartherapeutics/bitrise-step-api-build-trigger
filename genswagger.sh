#!/bin/sh
# Regenerates the swagger client using the latest definition file.
# Note: May cause build breakage if the definition has changed.
set -o errexit
set -o nounset
set -o pipefail

checkForUtility() {
  local UTIL="$1"
  hash "$UTIL" 2>/dev/null || { echo >&2 "Missing $UTIL utility. Please install $UTIL from brew or your package manager of choice."; exit 1; }
}

checkForUtility curl
checkForUtility jq
checkForUtility swagger

set -x
# Monkey-patch swagger definition to remove 200 and include 201 response to work around strict checks in generated client
curl 'https://api-docs.bitrise.io/docs/swagger.json' | jq 'del(.paths."/apps/{app-slug}/builds".post.responses."200") | .paths."/apps/{app-slug}/builds".post.responses."201" = {"description":"Created","schema":{"type":"object","$ref":"#/definitions/v0.BuildTriggerRespModel"}}' > swagger.json
rm -Rf models operations client
swagger generate client
