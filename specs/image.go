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
