/*
   Copyright Farcloser.

   Licensed under the Apache License, Version 2.0 (the "License");
   you may not use this file except in compliance with the License.
   You may obtain a copy of the License at

       http://www.apache.org/licenses/LICENSE-2.0

   Unless required by applicable law or agreed to in writing, software
   distributed under the License is distributed on an "AS IS" BASIS,
   WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
   See the License for the specific language governing permissions and
   limitations under the License.
*/

package reference_test

import (
	"testing"

	"github.com/opencontainers/go-digest"
	"gotest.tools/v3/assert"

	"go.farcloser.world/containers/reference"
)

func TestReference(t *testing.T) {
	t.Parallel()

	//nolint:lll
	needles := map[string]struct {
		Error         string
		String        string
		Suggested     string
		FamiliarName  string
		FamiliarMatch map[string]bool
		Protocol      reference.Protocol
		Digest        digest.Digest
		Path          string
		Domain        string
		Tag           string
		ExplicitTag   string
	}{
		"": {
			Error: "invalid reference format",
		},
		"∞": {
			Error: "invalid reference format",
		},
		"abcd:∞": {
			Error: "invalid reference format",
		},
		"abcd@sha256:∞": {
			Error: "invalid reference format",
		},
		"abcd@∞": {
			Error: "invalid reference format",
		},
		"abcd:foo@sha256:∞": {
			Error: "invalid reference format",
		},
		"abcd:foo@∞": {
			Error: "invalid reference format",
		},
		"sha256:whatever": {
			Error:        "",
			String:       "docker.io/library/sha256:whatever",
			Suggested:    "sha256-abcde",
			FamiliarName: "sha256",
			FamiliarMatch: map[string]bool{
				"*a*":                      true,
				"?ha25?":                   true,
				"[s-z]ha25[0-9]":           true,
				"[^a]ha25[^a-z]":           true,
				"*6:whatever":              true,
				"docker.io/library/sha256": false,
			},
			Protocol:    "",
			Digest:      "",
			Path:        "library/sha256",
			Domain:      "docker.io",
			Tag:         "whatever",
			ExplicitTag: "whatever",
		},
		"sha256:4b826db5f1f14d1db0b560304f189d4b17798ddce2278b7822c9d32313fe3f50": {
			Error:        "",
			String:       "sha256:4b826db5f1f14d1db0b560304f189d4b17798ddce2278b7822c9d32313fe3f50",
			Suggested:    "untitled-abcde",
			FamiliarName: "",
			Protocol:     "",
			Digest:       "sha256:4b826db5f1f14d1db0b560304f189d4b17798ddce2278b7822c9d32313fe3f50",
			Path:         "",
			Domain:       "",
			Tag:          "",
		},
		"4b826db5f1f14d1db0b560304f189d4b17798ddce2278b7822c9d32313fe3f50": {
			Error:        "",
			String:       "sha256:4b826db5f1f14d1db0b560304f189d4b17798ddce2278b7822c9d32313fe3f50",
			Suggested:    "untitled-abcde",
			FamiliarName: "",
			Protocol:     "",
			Digest:       "sha256:4b826db5f1f14d1db0b560304f189d4b17798ddce2278b7822c9d32313fe3f50",
			Path:         "",
			Domain:       "",
			Tag:          "",
		},
		"image_name": {
			Error:        "",
			String:       "docker.io/library/image_name:latest",
			Suggested:    "image_name-abcde",
			FamiliarName: "image_name",
			Protocol:     "",
			Digest:       "",
			Path:         "library/image_name",
			Domain:       "docker.io",
			Tag:          "latest",
			ExplicitTag:  "",
		},
		"library/image_name": {
			Error:        "",
			String:       "docker.io/library/image_name:latest",
			Suggested:    "image_name-abcde",
			FamiliarName: "image_name",
			Protocol:     "",
			Digest:       "",
			Path:         "library/image_name",
			Domain:       "docker.io",
			Tag:          "latest",
			ExplicitTag:  "",
		},
		"something/image_name": {
			Error:        "",
			String:       "docker.io/something/image_name:latest",
			Suggested:    "image_name-abcde",
			FamiliarName: "something/image_name",
			Protocol:     "",
			Digest:       "",
			Path:         "something/image_name",
			Domain:       "docker.io",
			Tag:          "latest",
			ExplicitTag:  "",
		},
		"docker.io/library/image_name": {
			Error:        "",
			String:       "docker.io/library/image_name:latest",
			Suggested:    "image_name-abcde",
			FamiliarName: "image_name",
			Protocol:     "",
			Digest:       "",
			Path:         "library/image_name",
			Domain:       "docker.io",
			Tag:          "latest",
			ExplicitTag:  "",
		},
		"image_name:latest": {
			Error:        "",
			String:       "docker.io/library/image_name:latest",
			Suggested:    "image_name-abcde",
			FamiliarName: "image_name",
			Protocol:     "",
			Digest:       "",
			Path:         "library/image_name",
			Domain:       "docker.io",
			Tag:          "latest",
			ExplicitTag:  "latest",
		},
		"image_name:foo": {
			Error:        "",
			String:       "docker.io/library/image_name:foo",
			Suggested:    "image_name-abcde",
			FamiliarName: "image_name",
			Protocol:     "",
			Digest:       "",
			Path:         "library/image_name",
			Domain:       "docker.io",
			Tag:          "foo",
			ExplicitTag:  "foo",
		},
		"image_name@sha256:4b826db5f1f14d1db0b560304f189d4b17798ddce2278b7822c9d32313fe3f50": {
			Error:        "",
			String:       "docker.io/library/image_name@sha256:4b826db5f1f14d1db0b560304f189d4b17798ddce2278b7822c9d32313fe3f50",
			Suggested:    "image_name-abcde",
			FamiliarName: "image_name",
			Protocol:     "",
			Digest:       "sha256:4b826db5f1f14d1db0b560304f189d4b17798ddce2278b7822c9d32313fe3f50",
			Path:         "library/image_name",
			Domain:       "docker.io",
			Tag:          "",
			ExplicitTag:  "",
		},
		"image_name:latest@sha256:4b826db5f1f14d1db0b560304f189d4b17798ddce2278b7822c9d32313fe3f50": {
			Error:        "",
			String:       "docker.io/library/image_name:latest@sha256:4b826db5f1f14d1db0b560304f189d4b17798ddce2278b7822c9d32313fe3f50",
			Suggested:    "image_name-abcde",
			FamiliarName: "image_name",
			Protocol:     "",
			Digest:       "sha256:4b826db5f1f14d1db0b560304f189d4b17798ddce2278b7822c9d32313fe3f50",
			Path:         "library/image_name",
			Domain:       "docker.io",
			Tag:          "latest",
			ExplicitTag:  "latest",
		},
		"ghcr.io:1234/image_name": {
			Error:        "",
			String:       "ghcr.io:1234/image_name:latest",
			Suggested:    "image_name-abcde",
			FamiliarName: "ghcr.io:1234/image_name",
			Protocol:     "",
			Digest:       "",
			Path:         "image_name",
			Domain:       "ghcr.io:1234",
			Tag:          "latest",
			ExplicitTag:  "",
		},
		"ghcr.io/sub_name/image_name": {
			Error:        "",
			String:       "ghcr.io/sub_name/image_name:latest",
			Suggested:    "image_name-abcde",
			FamiliarName: "ghcr.io/sub_name/image_name",
			Protocol:     "",
			Digest:       "",
			Path:         "sub_name/image_name",
			Domain:       "ghcr.io",
			Tag:          "latest",
			ExplicitTag:  "",
		},
		"oci-archive:///tmp/build/saved-image.tar": {
			Error: "image must be loaded from archive before parsing image reference",
		},
	}

	for index, test := range needles {
		parsed, err := reference.Parse(index)
		if test.Error != "" || err != nil {
			assert.Error(t, err, test.Error)

			continue
		}

		assert.Equal(t, parsed.String(), test.String, index)
		assert.Equal(t, parsed.SuggestContainerName("abcdefghij"), test.Suggested, index)
		assert.Equal(t, parsed.FamiliarName(), test.FamiliarName, index)

		for needle, result := range test.FamiliarMatch {
			res, err := parsed.FamiliarMatch(needle)
			assert.NilError(t, err)
			assert.Equal(t, res, result, index)
		}

		assert.Equal(t, parsed.Protocol, test.Protocol, index)
		assert.Equal(t, parsed.Digest, test.Digest, index)
		assert.Equal(t, parsed.Path, test.Path, index)
		assert.Equal(t, parsed.Domain, test.Domain, index)
		assert.Equal(t, parsed.Tag, test.Tag, index)
		assert.Equal(t, parsed.ExplicitTag, test.ExplicitTag, index)
	}
}
