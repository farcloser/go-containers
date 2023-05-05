package digest

import up "github.com/opencontainers/go-digest"

type Digest = up.Digest

func FromBytes(p []byte) Digest {
	return up.FromBytes(p)
}
