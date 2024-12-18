package specs

import (
	runtime "github.com/opencontainers/runtime-spec/specs-go"
)

type (
	Spec           = runtime.Spec
	Linux          = runtime.Linux
	LinuxResources = runtime.LinuxResources
	LinuxBlockIO   = runtime.LinuxBlockIO
	LinuxCPU       = runtime.LinuxCPU
	LinuxMemory    = runtime.LinuxMemory
	LinuxPids      = runtime.LinuxPids
)
