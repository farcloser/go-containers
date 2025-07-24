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

// Version returns the cgroup version in use on the system.
func Version() SystemVersion {
	return NoVersion
}

// DefaultManager returns the default cgroup manager based on the system configuration.
func DefaultManager() Manager {
	return NoManager
}

// DefaultMode returns the default cgroup mode.
func DefaultMode() Mode {
	return NoNsMode
}

// AvailableManagers returns a list of available cgroup managers.
func AvailableManagers() []Manager {
	return []Manager{}
}

// AvailableModes returns a list of available cgroup modes.
func AvailableModes() []Mode {
	return []Mode{}
}
