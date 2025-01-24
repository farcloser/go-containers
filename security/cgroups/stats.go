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

import "github.com/containerd/cgroups/v3/cgroup2/stats"

type Metrics = stats.Metrics

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
