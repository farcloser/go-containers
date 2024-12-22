package cgroups

import (
	"os"
	"path"
	"strings"

	"github.com/containerd/cgroups/v3"
	"github.com/containerd/cgroups/v3/cgroup2"
	"github.com/moby/sys/userns"
)

type Controller string

const (
	memoryController Controller = "memory"
	cpuController    Controller = "cpu"
	ioController     Controller = "io"
	cpuSetController Controller = "cpuset"
	pidsController   Controller = "pids"

	systemdPath            = "/run/systemd/system"
	cgroupNSPath           = "/proc/self/ns/cgroup"
	procSelfCGroupPath     = "/proc/self/cgroup"
	memorySwapMaxFile      = "memory.swap.max"
	cpuSetCPUEffectiveFile = "cpuset.cpus.effective"
	cpuSetMemEffectiveFile = "cpuset.mems.effective"
)

func Version() SystemVersion {
	if cgroups.Mode() == cgroups.Unified {
		return Version2
	}

	return Version1
}

func DefaultManager() Manager {
	if Version() == Version2 && isSystemdAvalailable() {
		return SystemdManager
	}

	return NoneManager
}

func DefaultMode() Mode {
	if Version() == Version2 {
		return PrivateNsMode
	}

	return NoNsMode
}

func AvailableManagers() []Manager {
	candidates := []Manager{NoneManager}
	if isSystemdAvalailable() {
		candidates = append(candidates, SystemdManager)
	}

	return candidates
}

func getSwapLimit() bool {
	_, unified, err := cgroups.ParseCgroupFileUnified(procSelfCGroupPath)
	if err != nil {
		return false
	}

	if unified == "" {
		return false
	}

	cGroupPath := path.Join(CgroupRoot, unified, memorySwapMaxFile)
	if _, err = os.Stat(cGroupPath); os.IsNotExist(err) {
		return false
	}

	return true
}

func GetCPUMemInfo(groupPath string) (string, string) {
	cpus, err := os.ReadFile(path.Join(CgroupRoot, groupPath, cpuSetCPUEffectiveFile))
	if err != nil {
		return "", ""
	}

	mems, err := os.ReadFile(path.Join(CgroupRoot, groupPath, cpuSetMemEffectiveFile))
	if err != nil {
		return "", ""
	}

	return strings.TrimSpace(string(cpus)), strings.TrimSpace(string(mems))
}

func New(pth string) (*Info, []error, error) {
	var warnings []error

	if pth == "" {
		pth = "/"
	}

	m, err := cgroup2.Load(pth)
	if err != nil {
		return nil, warnings, err
	}

	controllers, err := m.Controllers()
	if err != nil {
		return nil, warnings, err
	}

	ctrls := make(map[string]struct{}, len(controllers))
	for _, c := range controllers {
		ctrls[c] = struct{}{}
	}

	info := &Info{}

	if _, ok := ctrls[string(memoryController)]; !ok {
		warnings = append(warnings, ErrNoMemoryController)
	} else {
		info.MemoryLimit = true
		info.SwapLimit = getSwapLimit()
		info.MemoryReservation = true
		info.OomKillDisable = false
		info.MemorySwappiness = false
		info.KernelMemory = false
		info.KernelMemoryTCP = false
	}

	if _, ok := ctrls[string(cpuController)]; !ok {
		warnings = append(warnings, ErrNoCPUController)
	} else {
		info.CPUShares = true
		info.CPUCfs = true
		info.CPURealtime = false
	}

	if _, ok := ctrls[string(ioController)]; !ok {
		warnings = append(warnings, ErrNoIoController)
	} else {
		info.BlkioWeight = true
		info.BlkioWeightDevice = true
		info.BlkioReadBpsDevice = true
		info.BlkioWriteBpsDevice = true
		info.BlkioReadIOpsDevice = true
		info.BlkioWriteIOpsDevice = true
	}

	if _, ok := ctrls[string(cpuSetController)]; !ok {
		warnings = append(warnings, ErrNoCPUSetController)
	} else {
		info.Cpuset = true
		info.Cpus, info.Mems = GetCPUMemInfo(pth)
	}

	if _, ok := ctrls[string(pidsController)]; !ok {
		warnings = append(warnings, ErrNoPidsController)
	} else {
		info.PidsLimit = true
	}

	info.CgroupDevicesEnabled = !userns.RunningInUserNS()

	if _, err = os.Stat(cgroupNSPath); !os.IsNotExist(err) {
		info.CgroupNamespaces = true
	}

	return info, warnings, nil
}

func isSystemdAvalailable() bool {
	fi, err := os.Lstat(systemdPath)
	if err != nil {
		return false
	}

	return fi.IsDir()
}
