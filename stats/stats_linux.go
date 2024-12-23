package stats

import (
	"bufio"
	"os"
	"strconv"
	"strings"
	"time"

	v2 "github.com/containerd/cgroups/v3/cgroup2/stats"

	"go.farcloser.world/containers/cgroups"
	"go.farcloser.world/containers/netlink"
)

const (
	microsecPerSecond = 1000
	percent           = 100.0
	kiloPerMega       = 1024
	procMemInfoPath   = "/proc/meminfo"
	memoryMaxLimit    = float64(^uint64(0))
)

func SetCgroup2StatsFields(previousStats *ContainerStats, metrics *v2.Metrics, links []netlink.Link) (Entry, error) {
	netRx, netTx := netlink.StatsForLinks(links)
	blkRead, blkWrite := cgroups.CalculateIO(metrics)
	mem := cgroups.CalculateMemUsage(metrics)

	memLimit := float64(metrics.GetMemory().GetUsageLimit())
	if memLimit == memoryMaxLimit {
		memLimit = getHostMemLimit()
	}

	memPercent := calculateMemPercent(memLimit, mem)
	pidsStatsCurrent := metrics.GetPids().GetCurrent()

	cpuPercent := calculateCgroup2CPUPercent(previousStats, metrics)

	return Entry{
		CPUPercentage:    cpuPercent,
		Memory:           mem,
		MemoryPercentage: memPercent,
		MemoryLimit:      memLimit,
		NetworkRx:        netRx,
		NetworkTx:        netTx,
		BlockRead:        float64(blkRead),
		BlockWrite:       float64(blkWrite),
		PidsCurrent:      pidsStatsCurrent,
	}, nil
}

func calculateMemPercent(limit float64, usedNo float64) float64 {
	// Limit will never be 0 unless the container is not running, and we haven't
	// got any data from cgroup
	if limit != 0 {
		return usedNo / limit * percent
	}

	return 0
}

func getHostMemLimit() float64 {
	file, err := os.Open(procMemInfoPath)
	if err != nil {
		return memoryMaxLimit
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		if strings.HasPrefix(scanner.Text(), "MemTotal:") {
			fields := strings.Fields(scanner.Text())
			if len(fields) > 1 {
				memKb, err := strconv.ParseUint(fields[1], 10, 64)
				if err == nil {
					return float64(memKb * kiloPerMega) // kB to bytes
				}
			}

			break
		}
	}

	return float64(^uint64(0))
}

// PercpuUsage is not supported in CgroupV2.
func calculateCgroup2CPUPercent(previousStats *ContainerStats, metrics *v2.Metrics) float64 {
	var (
		cpuPercent = 0.0
		cpu        = metrics.GetCPU()
		// calculate the change for the cpu usage of the container in between readings
		cpuDelta = float64(cpu.GetUsageUsec()*microsecPerSecond) - float64(previousStats.Cgroup2CPU)
		// calculate the change for the entire system between readings
		_ = float64(cpu.GetSystemUsec()*microsecPerSecond) - float64(previousStats.Cgroup2System)
		// time duration
		timeDelta = time.Since(previousStats.Time)
	)

	if cpuDelta > 0.0 {
		cpuPercent = cpuDelta / float64(timeDelta.Nanoseconds()) * percent
	}

	return cpuPercent
}
