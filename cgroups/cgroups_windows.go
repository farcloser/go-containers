package cgroups

// Cgroups not supported on Windows.
func Version() int {
	return 0
}
