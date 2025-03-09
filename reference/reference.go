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

package reference

import (
	"errors"
	"path"
	"strings"

	"github.com/distribution/reference"
	"github.com/opencontainers/go-digest"
)

type Protocol string

const shortIDLength = 5

var ErrLoadOCIArchiveRequired = errors.New("image must be loaded from archive before parsing image reference")

type ImageReference struct {
	Protocol    Protocol
	Digest      digest.Digest
	Tag         string
	ExplicitTag string
	Path        string
	Domain      string

	nn reference.Reference
}

func (ir *ImageReference) Name() string {
	ret := ir.Domain
	if ret != "" {
		ret += "/"
	}

	ret += ir.Path

	return ret
}

func (ir *ImageReference) FamiliarName() string {
	if ir.Protocol != "" && ir.Domain == "" {
		return ir.Path
	}

	if ir.nn != nil {
		if v, ok := ir.nn.(reference.Named); ok {
			return reference.FamiliarName(v)
		}
	}

	return ""
}

func (ir *ImageReference) FamiliarMatch(pattern string) (bool, error) {
	if ir.nn != nil {
		return reference.FamiliarMatch(pattern, ir.nn)
	}

	return false, nil
}

func (ir *ImageReference) String() string {
	if ir.Protocol != "" && ir.Domain == "" {
		return ir.Path
	}

	if ir.Path == "" && ir.Digest != "" {
		return ir.Digest.String()
	}

	if ir.nn != nil {
		return ir.nn.String()
	}

	return ""
}

func (ir *ImageReference) SuggestContainerName(suffix string) string {
	name := "untitled"
	if ir.Protocol != "" && ir.Domain == "" {
		name = string(ir.Protocol) + "-" + ir.String()[:shortIDLength]
	} else if ir.Path != "" {
		name = path.Base(ir.Path)
	}

	return name + "-" + suffix[:5]
}

func Parse(rawRef string) (*ImageReference, error) {
	imageRef := &ImageReference{}

	if strings.HasPrefix(rawRef, "oci-archive://") {
		// The image must be loaded from the specified archive path first
		// before parsing the image reference specified in its OCI image manifest.
		return nil, ErrLoadOCIArchiveRequired
	}

	if dgst, err := digest.Parse(rawRef); err == nil {
		imageRef.Digest = dgst

		return imageRef, nil
	} else if dgst, err := digest.Parse("sha256:" + rawRef); err == nil {
		imageRef.Digest = dgst

		return imageRef, nil
	}

	var err error

	imageRef.nn, err = reference.ParseNormalizedNamed(rawRef)
	if err != nil {
		return imageRef, err
	}

	if tg, ok := imageRef.nn.(reference.Tagged); ok {
		imageRef.ExplicitTag = tg.Tag()
	}

	if tg, ok := imageRef.nn.(reference.Named); ok {
		imageRef.nn = reference.TagNameOnly(tg)
		imageRef.Domain = reference.Domain(tg)
		imageRef.Path = reference.Path(tg)
	}

	if tg, ok := imageRef.nn.(reference.Tagged); ok {
		imageRef.Tag = tg.Tag()
	}

	if tg, ok := imageRef.nn.(reference.Digested); ok {
		imageRef.Digest = tg.Digest()
	}

	return imageRef, nil
}
