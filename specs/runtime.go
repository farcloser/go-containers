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
	LinuxResources    = runtime.LinuxResources
	LinuxBlockIO      = runtime.LinuxBlockIO
	LinuxCPU          = runtime.LinuxCPU
	LinuxMemory       = runtime.LinuxMemory
	LinuxPids         = runtime.LinuxPids
	LinuxCapabilities = runtime.LinuxCapabilities
	LinuxNamespace    = runtime.LinuxNamespace

	POSIXRlimit = runtime.POSIXRlimit
)
