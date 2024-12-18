package specs

import (
	"github.com/opencontainers/image-spec/identity"
	root "github.com/opencontainers/image-spec/specs-go"
	images "github.com/opencontainers/image-spec/specs-go/v1"

	"go.farcloser.world/containers/digest"
)

type (
	Platform    = images.Platform
	Descriptor  = images.Descriptor
	Image       = images.Image
	Index       = images.Index
	Manifest    = images.Manifest
	ImageConfig = images.ImageConfig
	History     = images.History
	ImageLayout = images.ImageLayout
	RootFS      = images.RootFS
	Versioned   = root.Versioned
)

const (
	MediaTypeImageManifest  = images.MediaTypeImageManifest
	MediaTypeImageConfig    = images.MediaTypeImageConfig
	MediaTypeImageLayerZstd = images.MediaTypeImageLayerZstd
	MediaTypeImageLayerGzip = images.MediaTypeImageLayerGzip
	MediaTypeImageIndex     = images.MediaTypeImageIndex
	MediaTypeImageLayer     = images.MediaTypeImageLayer
)

func ChainID(dgsts []digest.Digest) digest.Digest {
	return identity.ChainID(dgsts)
}
