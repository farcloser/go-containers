package cgroups

import (
	"github.com/containerd/cgroups/v3"
)

const (
	Version1 = 1
	Version2 = 2
)

func Version() int {
	if cgroups.Mode() == cgroups.Unified {
		return Version2
	}

	return Version1
}
