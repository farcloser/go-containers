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
