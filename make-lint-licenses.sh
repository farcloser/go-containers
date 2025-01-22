#!/usr/bin/env bash
# shellcheck disable=SC2034,SC2015
set -o errexit -o errtrace -o functrace -o nounset -o pipefail

# FIXME: go-licenses cannot find LICENSE from root of repo when submodule is imported:
# https://github.com/google/go-licenses/issues/186
# This is impacting gotest.tools
# go-licenses is also really broken right now wrt to stdlib: https://github.com/google/go-licenses/issues/244
# workaround taken from the awesome folks at Pulumi: https://github.com/pulumi/license-check-action/pull/3
go-licenses check --include_tests --allowed_licenses=Apache-2.0,BSD-2-Clause,BSD-3-Clause,MIT \
	  --ignore gotest.tools \
	  --ignore "$(go list std | awk 'NR > 1 { printf(",") } { printf("%s",$0) } END { print "" }')" \
	  ./...

printf "WARNING: you need to manually verify licenses for:\n- gotest.tools\n"
