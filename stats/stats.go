package stats

import (
	"errors"
	"sync"
	"time"
)

var (
	ErrFailedConversion = errors.New("cannot convert metric data to cgroups.Metrics")
	ErrEmptyMetrics     = errors.New("nothing in provided metric")
)

// Entry represents the statistics data collected from a container.
type Entry struct {
	Name             string
	ID               string
	CPUPercentage    float64
	Memory           float64
	MemoryLimit      float64
	MemoryPercentage float64
	NetworkRx        float64
	NetworkTx        float64
	BlockRead        float64
	BlockWrite       float64
	PidsCurrent      uint64
	IsInvalid        bool
}

// ContainerStats represents the runtime container stats.
type ContainerStats struct {
	Time          time.Time
	Cgroup2CPU    uint64
	Cgroup2System uint64
}

// Stats represents an entity to store containers statistics synchronously.
type Stats struct {
	mutex sync.RWMutex
	Entry
	err error
}

// NewStats is from
// https://github.com/docker/cli/blob/3fb4fb83dfb5db0c0753a8316f21aea54dab32c5/cli/command/container/formatter_stats.go#L113-L116
//
//nolint:lll
func NewStats(containerID string) *Stats {
	return &Stats{Entry: Entry{ID: containerID}}
}

func (cs *Stats) SetStatistics(s Entry) {
	cs.mutex.Lock()
	defer cs.mutex.Unlock()

	cs.Entry = s
}

func (cs *Stats) GetStatistics() Entry {
	cs.mutex.Lock()
	defer cs.mutex.Unlock()

	return cs.Entry
}

func (cs *Stats) GetError() error {
	cs.mutex.Lock()
	defer cs.mutex.Unlock()

	return cs.err
}

func (cs *Stats) SetErrorAndReset(err error) {
	cs.mutex.Lock()
	defer cs.mutex.Unlock()

	cs.CPUPercentage = 0
	cs.Memory = 0
	cs.MemoryPercentage = 0
	cs.MemoryLimit = 0
	cs.NetworkRx = 0
	cs.NetworkTx = 0
	cs.BlockRead = 0
	cs.BlockWrite = 0
	cs.PidsCurrent = 0
	cs.err = err
	cs.IsInvalid = true
}

func (cs *Stats) SetError(err error) {
	cs.mutex.Lock()
	defer cs.mutex.Unlock()

	cs.err = err

	if err != nil {
		cs.IsInvalid = true
	}
}
