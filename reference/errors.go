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

import "errors"

var (
	// ErrInvalidImageReference indicates that the image reference is invalid.
	ErrInvalidImageReference = errors.New("invalid image reference")
	// ErrInvalidPattern indicates that the pattern used to parse the image reference is invalid.
	ErrInvalidPattern = errors.New("invalid pattern")
	// ErrLoadOCIArchiveRequired indicates that the image must be loaded from an OCI archive.
	ErrLoadOCIArchiveRequired = errors.New("image must be loaded from archive before parsing image reference")
)
