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
	runtime "github.com/opencontainers/runtime-spec/specs-go"
)

const (
	CgroupNamespace  = runtime.CgroupNamespace
	IPCNamespace     = runtime.IPCNamespace
	PIDNamespace     = runtime.PIDNamespace
	UTSNamespace     = runtime.UTSNamespace
	NetworkNamespace = runtime.NetworkNamespace
)

type (
	Spec    = runtime.Spec
	Root    = runtime.Root
	State   = runtime.State
	Mount   = runtime.Mount
	Box     = runtime.Box
	Process = runtime.Process
	Hook    = runtime.Hook
	Hooks   = runtime.Hooks

	Linux             = runtime.Linux
	Windows           = runtime.Windows
	LinuxResources    = runtime.LinuxResources
	LinuxBlockIO      = runtime.LinuxBlockIO
	LinuxCPU          = runtime.LinuxCPU
	LinuxMemory       = runtime.LinuxMemory
	LinuxPids         = runtime.LinuxPids
	LinuxCapabilities = runtime.LinuxCapabilities
	LinuxNamespace    = runtime.LinuxNamespace

	POSIXRlimit = runtime.POSIXRlimit
)
