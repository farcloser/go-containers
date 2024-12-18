package specs

import (
	"github.com/opencontainers/image-spec/identity"
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
