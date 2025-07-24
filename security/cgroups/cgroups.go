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

import (
	"errors"
)

// SystemVersion represents the version of the cgroup system in use.
type SystemVersion int

type (
	// Manager represents the cgroup manager in use.
	Manager string
	// Mode represents the namespace mode for cgroups.
	Mode string
)

const (
	// NoVersion indicates that no specific cgroup version is set.
	NoVersion SystemVersion = 0
	// Version1 indicates that cgroup v1 is in use.
	Version1 SystemVersion = 1
	// Version2 indicates that cgroup v2 is in use.
	Version2 SystemVersion = 2

	// NoManager indicates that no specific cgroup manager is set.
	NoManager Manager = ""
	// NoneManager indicates that no cgroup manager is used.
	NoneManager Manager = "none"
	// SystemdManager indicates that systemd is used as the cgroup manager.
	SystemdManager Manager = "systemd"

	// NoNsMode indicates that no specific namespace mode is set.
	NoNsMode Mode = ""
	// HostNsMode indicates that the host namespace mode is used.
	HostNsMode Mode = "host"
	// PrivateNsMode indicates that the private namespace mode is used.
	PrivateNsMode Mode = "private"
)

var (
	// ErrNoMemoryController is returned when the system does not support memory controller.
	ErrNoMemoryController = errors.New("no systemd memory controller found")
	// ErrNoCPUController is returned when the system does not support CPU controller.
	ErrNoCPUController = errors.New("no systemd cpu controller found")
	// ErrNoIoController is returned when the system does not support IO controller.
	ErrNoIoController = errors.New("no systemd io controller found")
	// ErrNoCPUSetController is returned when the system does not support CPU set controller.
	ErrNoCPUSetController = errors.New("no systemd cpuset controller found")
	// ErrNoPidsController is returned when the system does not support PIDs controller.
	ErrNoPidsController = errors.New("no systemd pids controller found")
)

// Info contains information about the cgroup system, including memory, CPU, block IO,.
type Info struct {
	memInfo
	cpuInfo
	blkioInfo
	cpuSetInfo
	pidsInfo
	genericInfo
}

type memInfo struct {
	// Whether memory limit is supported or not
	MemoryLimit bool

	// Whether swap limit is supported or not
	SwapLimit bool

	// Whether soft limit is supported or not
	MemoryReservation bool

	// Whether OOM killer disable is supported or not
	OomKillDisable bool

	// Whether memory swappiness is supported or not
	MemorySwappiness bool

	// Whether kernel memory limit is supported or not. This option is used to
	// detect support for kernel-memory limits on API < v1.42. Kernel memory
	// limit (`kmem.limit_in_bytes`) is not supported on cgroups v2, and has been
	// removed in kernel 5.4.
	KernelMemory bool

	// Whether kernel memory TCP limit is supported or not. Kernel memory TCP
	// limit (`memory.kmem.tcp.limit_in_bytes`) is not supported on cgroups v2.
	KernelMemoryTCP bool
}

type cpuInfo struct {
	// Whether CPU shares is supported or not
	CPUShares bool

	// Whether CPU CFS (Completely Fair Scheduler) is supported
	CPUCfs bool

	// Whether CPU real-time scheduler is supported
	CPURealtime bool
}

type blkioInfo struct {
	// Whether Block IO weight is supported or not
	BlkioWeight bool

	// Whether Block IO weight_device is supported or not
	BlkioWeightDevice bool

	// Whether Block IO read limit in bytes per second is supported or not
	BlkioReadBpsDevice bool

	// Whether Block IO write limit in bytes per second is supported or not
	BlkioWriteBpsDevice bool

	// Whether Block IO read limit in IO per second is supported or not
	BlkioReadIOpsDevice bool

	// Whether Block IO write limit in IO per second is supported or not
	BlkioWriteIOpsDevice bool
}

type cpuSetInfo struct {
	// Whether Cpuset is supported or not
	Cpuset bool

	// Available Cpuset's cpus
	Cpus string

	// Available Cpuset's memory nodes
	Mems string
}

type pidsInfo struct {
	// Whether Pids Limit is supported or not
	PidsLimit bool
}

type genericInfo struct {
	// Whether the cgroup has the mountpoint of "devices" or not
	CgroupDevicesEnabled bool
	// Whether the kernel supports cgroup namespaces or not
	CgroupNamespaces bool
}
