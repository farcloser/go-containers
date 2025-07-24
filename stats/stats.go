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

package stats

import (
	"errors"
	"sync"
	"time"
)

var (
	// ErrFailedConversion is returned when the conversion from metric data to cgroups.Metrics fails.
	ErrFailedConversion = errors.New("cannot convert metric data to cgroups.Metrics")
	// ErrEmptyMetrics is returned when the provided metric data is empty.
	ErrEmptyMetrics = errors.New("nothing in provided metric")
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
func NewStats(containerID string) *Stats {
	return &Stats{Entry: Entry{ID: containerID}}
}

// SetStatistics sets the statistics for the container stats synchronously.
func (cs *Stats) SetStatistics(s Entry) {
	cs.mutex.Lock()
	defer cs.mutex.Unlock()

	cs.Entry = s
}

// GetStatistics retrieves the statistics for the container stats synchronously.
func (cs *Stats) GetStatistics() Entry {
	cs.mutex.Lock()
	defer cs.mutex.Unlock()

	return cs.Entry
}

// GetError retrieves the error associated with the stats synchronously.
func (cs *Stats) GetError() error {
	cs.mutex.Lock()
	defer cs.mutex.Unlock()

	return cs.err
}

// SetErrorAndReset sets the error and resets the statistics fields to their zero values.
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

// SetError sets the error and marks the stats as invalid.
func (cs *Stats) SetError(err error) {
	cs.mutex.Lock()
	defer cs.mutex.Unlock()

	cs.err = err

	if err != nil {
		cs.IsInvalid = true
	}
}
