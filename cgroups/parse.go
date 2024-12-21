// Forked from https://github.com/moby/moby/blob/v27.4.1/pkg/parsers/parsers.go

// Package parsers provides helper functions to parse and validate different type
// of string. It can be hosts, unix addresses, tcp addresses, filters, kernel
// operating system versions.
package cgroups // import "github.com/docker/docker/pkg/parsers"

import (
	"errors"
	"fmt"
	"strconv"
	"strings"
)

var (
	errValueOutOfRange = errors.New("value of out range")
	errInvalidFormat   = errors.New("invalid format")
)

// ParseUintListMaximum parses and validates the specified string as the value
// found in some cgroup file (e.g. `cpuset.cpus`, `cpuset.mems`), which could be
// one of the formats below. Note that duplicates are actually allowed in the
// input string. It returns a `map[int]bool` with available elements from `val`
// set to `true`. Values larger than `maximum` cause an error if max is non zero,
// in order to stop the map becoming excessively large.
// Supported formats:
//
//	7
//	1-6
//	0,3-4,7,8-10
//	0-0,0,1-7
//	03,1-3      <- this is going to get parsed as [1,2,3]
//	3,2,1
//	0-2,3,1
func ParseUintListMaximum(val string, maximum int) (map[int]bool, error) {
	return parseUintList(val, maximum)
}

// ParseUintList parses and validates the specified string as the value
// found in some cgroup file (e.g. `cpuset.cpus`, `cpuset.mems`), which could be
// one of the formats below. Note that duplicates are actually allowed in the
// input string. It returns a `map[int]bool` with available elements from `val`
// set to `true`.
// Supported formats:
//
//	7
//	1-6
//	0,3-4,7,8-10
//	0-0,0,1-7
//	03,1-3      <- this is going to get parsed as [1,2,3]
//	3,2,1
//	0-2,3,1
func ParseUintList(val string) (map[int]bool, error) {
	return parseUintList(val, 0)
}

func parseUintList(val string, maximum int) (map[int]bool, error) {
	if val == "" {
		return map[int]bool{}, nil
	}

	availableInts := make(map[int]bool)
	split := strings.Split(val, ",")
	errInvalid := fmt.Errorf("%w: %s", errInvalidFormat, val)

	for _, value := range split {
		if !strings.Contains(value, "-") { //nolint:nestif
			intValue, err := strconv.Atoi(value)
			if err != nil {
				return nil, errInvalid
			}

			if maximum != 0 && intValue > maximum {
				return nil, fmt.Errorf("%w, maximum is %d", errValueOutOfRange, maximum)
			}

			availableInts[intValue] = true
		} else {
			minS, maxS, _ := strings.Cut(value, "-")

			minInt, err := strconv.Atoi(minS)
			if err != nil {
				return nil, errInvalid
			}

			maxInt, err := strconv.Atoi(maxS)
			if err != nil {
				return nil, errInvalid
			}

			if maxInt < minInt {
				return nil, errInvalid
			}

			if maximum != 0 && maxInt > maximum {
				return nil, fmt.Errorf("%w, maximum is %d", errValueOutOfRange, maximum)
			}

			for i := minInt; i <= maxInt; i++ {
				availableInts[i] = true
			}
		}
	}

	return availableInts, nil
}
