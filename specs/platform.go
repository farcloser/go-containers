package specs

import (
	v1 "github.com/opencontainers/image-spec/specs-go/v1"
)

type (
	Platform    = v1.Platform
	Descriptor  = v1.Descriptor
	Image       = v1.Image
	Index       = v1.Index
	Manifest    = v1.Manifest
	ImageConfig = v1.ImageConfig
	History     = v1.History
	ImageLayout = v1.ImageLayout
	RootFS      = v1.RootFS
)

const (
	MediaTypeImageManifest  = v1.MediaTypeImageManifest
	MediaTypeImageConfig    = v1.MediaTypeImageConfig
	MediaTypeImageLayerZstd = v1.MediaTypeImageLayerZstd
	MediaTypeImageLayerGzip = v1.MediaTypeImageLayerGzip
	MediaTypeImageIndex     = v1.MediaTypeImageIndex
	MediaTypeImageLayer     = v1.MediaTypeImageLayer
)
