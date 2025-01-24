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

var (
	// ErrDigestInvalidFormat returned when digest format invalid.
	ErrDigestInvalidFormat = up.ErrDigestInvalidFormat

	// ErrDigestInvalidLength returned when digest has invalid length.
	ErrDigestInvalidLength = up.ErrDigestInvalidLength

	// ErrDigestUnsupported returned when the digest algorithm is unsupported.
	ErrDigestUnsupported = up.ErrDigestUnsupported
)

type Digest = up.Digest

func FromBytes(p []byte) Digest {
	return up.FromBytes(p)
}

func FromString(s string) Digest {
	return FromBytes([]byte(s))
}

func Parse(s string) (Digest, error) {
	return up.Parse(s)
}
