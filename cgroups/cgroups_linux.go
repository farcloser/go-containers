package cgroups

import (
	"github.com/containerd/cgroups/v3"
)

const (
	version1 = 1
	version2 = 2
)

func Version() int {
	if cgroups.Mode() == cgroups.Unified {
		return version2
	}

	return version1
}
