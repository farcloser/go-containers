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

package sysinfo

import "go.farcloser.world/containers/cgroups"

type SysInfo struct {
	cgroups.Info

	// Whether the kernel supports AppArmor or not
	AppArmor bool

	// Whether the kernel supports Seccomp or not
	Seccomp bool

	// Whether IPv4 forwarding is supported or not, if this was disabled, networking will not work
	IPv4ForwardingDisabled bool

	// Warnings contains a slice of warnings that occurred  while collecting
	// system information. These warnings are intended to be informational
	// messages for the user, and can either be logged or returned to the
	// client; they are not intended to be parsed / used for other purposes,
	// and do not have a fixed format.
	Warnings []error
}
