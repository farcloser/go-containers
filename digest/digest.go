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

package digest

import (
	up "github.com/opencontainers/go-digest"
)

// Digest allows simple protection of hex formatted digest strings, prefixed
// by their algorithm. Strings of type Digest have some guarantee of being in
// the correct format and it provides quick access to the components of a
// digest string.
type Digest = up.Digest

// FromBytes creates a Digest from a byte slice.
func FromBytes(p []byte) Digest {
	return up.FromBytes(p)
}

// FromString creates a Digest from a string.
func FromString(s string) Digest {
	return FromBytes([]byte(s))
}

// Parse parses a digest string and returns a Digest.
func Parse(s string) (Digest, error) {
	//nolint:wrapcheck
	return up.Parse(s)
}
