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

package specs

import (
	"github.com/opencontainers/go-digest"
	"github.com/opencontainers/image-spec/identity"
	root "github.com/opencontainers/image-spec/specs-go"
	images "github.com/opencontainers/image-spec/specs-go/v1"
)

type (
	// Platform describes the platform which the image in the manifest runs on.
	Platform = images.Platform
	// Descriptor describes the disposition of targeted content.
	Descriptor = images.Descriptor
	// Image is the image manifest.
	Image = images.Image
	// Index is the image index.
	Index = images.Index
	// Manifest is the image manifest.
	Manifest = images.Manifest
	// ImageConfig is the image configuration.
	ImageConfig = images.ImageConfig
	// History is the history of the image.
	History = images.History
	// ImageLayout is the image layout.
	ImageLayout = images.ImageLayout
	// RootFS is the root filesystem of the image.
	RootFS = images.RootFS
	// Versioned is the versioned image specification.
	Versioned = root.Versioned
)

const (
	// MediaTypeDescriptor specifies the media type for a content descriptor.
	MediaTypeDescriptor = images.MediaTypeDescriptor
	// MediaTypeLayoutHeader specifies the media type for the oci-layout.
	MediaTypeLayoutHeader = images.MediaTypeLayoutHeader
	// MediaTypeEmptyJSON specifies the media type for an unused blob containing the value "{}".
	MediaTypeEmptyJSON = images.MediaTypeEmptyJSON
	// MediaTypeImageManifest specifies the media type for an image manifest.
	MediaTypeImageManifest = images.MediaTypeImageManifest
	// MediaTypeImageConfig specifies the media type for the image configuration.
	MediaTypeImageConfig = images.MediaTypeImageConfig
	// MediaTypeImageLayerZstd is the media type used for zstd compressed
	// layers referenced by the manifest.
	MediaTypeImageLayerZstd = images.MediaTypeImageLayerZstd
	// MediaTypeImageLayerGzip is the media type used for gzipped layers
	// referenced by the manifest.
	MediaTypeImageLayerGzip = images.MediaTypeImageLayerGzip
	// MediaTypeImageIndex specifies the media type for an image index.
	MediaTypeImageIndex = images.MediaTypeImageIndex
	// MediaTypeImageLayer is the media type used for layers referenced by the manifest.
	MediaTypeImageLayer = images.MediaTypeImageLayer
)

// ChainID takes a slice of digests and returns the ChainID corresponding to
// the last entry.
func ChainID(dgsts []digest.Digest) digest.Digest {
	return identity.ChainID(dgsts)
}
