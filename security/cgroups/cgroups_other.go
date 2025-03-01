//go:build !linux

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

package cgroups

func Version() SystemVersion {
	return NoVersion
}

func DefaultManager() Manager {
	return NoManager
}

func DefaultMode() Mode {
	return NoNsMode
}

func AvailableManagers() []Manager {
	return []Manager{}
}

func AvailableModes() []Mode {
	return []Mode{}
}
