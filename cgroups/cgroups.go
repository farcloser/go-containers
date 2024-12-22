package cgroups

import (
	"errors"
)

type SystemVersion int

type (
	Manager string
	Mode    string
)

const (
	NoVersion SystemVersion = 0
	Version1  SystemVersion = 1
	Version2  SystemVersion = 2

	SystemdManager Manager = "systemd"
	NoneManager    Manager = "none"
	NoManager      Manager = ""

	PrivateNsMode Mode = "private"
	HostNsMode    Mode = "host"
	NoNsMode      Mode = ""

	CgroupRoot = "/sys/fs/cgroup"
)

var (
	ErrNoMemoryController = errors.New("no systemd memory controller found")
	ErrNoCPUController    = errors.New("no systemd cpu controller found")
	ErrNoIoController     = errors.New("no systemd io controller found")
	ErrNoCPUSetController = errors.New("no systemd cpuset controller found")
	ErrNoPidsController   = errors.New("no systemd pids controller found")
)

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

// IsCpusetCpusAvailable returns `true` if the provided string set is contained
// in cgroup's cpuset.cpus set, `false` otherwise.
// If error is not nil a parsing error occurred.
func (c cpuSetInfo) IsCpusetCpusAvailable(provided string) (bool, error) {
	return isCpusetListAvailable(provided, c.Cpus)
}

// IsCpusetMemsAvailable returns `true` if the provided string set is contained
// in cgroup's cpuset.mems set, `false` otherwise.
// If error is not nil a parsing error occurred.
func (c cpuSetInfo) IsCpusetMemsAvailable(provided string) (bool, error) {
	return isCpusetListAvailable(provided, c.Mems)
}

func isCpusetListAvailable(provided, available string) (bool, error) {
	parsedAvailable, err := ParseUintList(available)
	if err != nil {
		return false, err
	}

	// 8192 is the normal maximum number of CPUs in Linux, so accept numbers up to this
	// or more if we actually have more CPUs.
	maxCPUs := 8192
	for m := range parsedAvailable {
		if m > maxCPUs {
			maxCPUs = m
		}
	}

	parsedProvided, err := ParseUintListMaximum(provided, maxCPUs)
	if err != nil {
		return false, err
	}

	for k := range parsedProvided {
		if !parsedAvailable[k] {
			return false, nil
		}
	}

	return true, nil
}