package cgroups

import "github.com/containerd/cgroups/v3/cgroup2/stats"

func CalculateMemUsage(metrics *stats.Metrics) float64 {
	usage := metrics.GetMemory().GetUsage()
	if v := metrics.GetMemory().GetInactiveFile(); v < usage {
		return float64(usage - v)
	}

	return float64(usage)
}

func CalculateIO(metrics *stats.Metrics) (uint64, uint64) {
	var ioRead, ioWrite uint64

	for _, iOEntry := range metrics.GetIo().GetUsage() {
		rios := iOEntry.GetRios()
		wios := iOEntry.GetWios()

		if rios == 0 && wios == 0 {
			continue
		}

		if rios != 0 {
			ioRead += iOEntry.GetRbytes()
		}

		if wios != 0 {
			ioWrite += iOEntry.GetWbytes()
		}
	}

	return ioRead, ioWrite
}
