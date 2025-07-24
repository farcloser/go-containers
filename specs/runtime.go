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
	// CgroupNamespace for isolating cgroup hierarchies.
	CgroupNamespace = runtime.CgroupNamespace
	// IPCNamespace for isolating System V IPC, POSIX message queues.
	IPCNamespace = runtime.IPCNamespace
	// PIDNamespace for isolating process IDs.
	PIDNamespace = runtime.PIDNamespace
	// UTSNamespace for isolating hostname and NIS domain name.
	UTSNamespace = runtime.UTSNamespace
	// NetworkNamespace for isolating network devices, stacks, ports, etc.
	NetworkNamespace = runtime.NetworkNamespace
	// MountNamespace for isolating mount points.
	MountNamespace = runtime.MountNamespace
	// UserNamespace for isolating user and group IDs.
	UserNamespace = runtime.UserNamespace
	// TimeNamespace for isolating the clocks.
	TimeNamespace = runtime.TimeNamespace
)

type (
	// Spec is the runtime spec for a container.
	Spec = runtime.Spec
	// Root is the root filesystem for a container.
	Root = runtime.Root
	// State is the state of a container.
	State = runtime.State
	// Mount is a mount point for a container.
	Mount = runtime.Mount
	// Box is a display console box for a container.
	Box = runtime.Box
	// Process is the process that runs in a container.
	Process = runtime.Process
	// Hook is a hook that can be run at various points in the container lifecycle.
	Hook = runtime.Hook
	// Hooks is a collection of hooks that can be run at various points in the container lifecycle.
	Hooks = runtime.Hooks

	// Linux is the Linux-specific configuration for a container.
	Linux = runtime.Linux
	// Windows is the Windows-specific configuration for a container.
	Windows = runtime.Windows
	// LinuxResources is the Linux-specific resources for a container.
	LinuxResources = runtime.LinuxResources
	// LinuxBlockIO is the Linux-specific block IO configuration for a container.
	LinuxBlockIO = runtime.LinuxBlockIO
	// LinuxCPU is the Linux-specific CPU configuration for a container.
	LinuxCPU = runtime.LinuxCPU
	// LinuxMemory is the Linux-specific memory configuration for a container.
	LinuxMemory = runtime.LinuxMemory
	// LinuxPids is the Linux-specific PIDs configuration for a container.
	LinuxPids = runtime.LinuxPids
	// LinuxCapabilities is the Linux-specific capabilities configuration for a container.
	LinuxCapabilities = runtime.LinuxCapabilities
	// LinuxNamespace is the Linux-specific namespace configuration for a container.
	LinuxNamespace = runtime.LinuxNamespace

	// POSIXRlimit is the POSIX resource limit configuration for a container.
	POSIXRlimit = runtime.POSIXRlimit
)
