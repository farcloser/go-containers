//go:build !linux

package cgroups

func Version() SystemVersion {
	return NoVersion
}

func DefaultManager() Manager {
	return NoManager
}

func DefaultMode() Mode {
	return NoNsMode
}

func AvailableManagers() []Manager {
	return []Manager{}
}

func AvailableModes() []Mode {
	return []Mode{}
}
