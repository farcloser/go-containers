package sysinfo

import (
	"os"
	"path/filepath"
	"strings"

	"github.com/containerd/containerd/v2/pkg/seccomp"

	"go.farcloser.world/containers/cgroups"
)

const (
	procSysNetPath = "/proc/sys/net"
	appArmorPath   = "/sys/kernel/security/apparmor"
)

func New(path string) (*SysInfo, []error, error) {
	if path == "" {
		path = "/"
	}

	info, warnings, err := cgroups.New(path)
	if err != nil {
		return nil, warnings, err
	}

	sysInfo := &SysInfo{
		Info: *info,
	}

	sysInfo.IPv4ForwardingDisabled = !readProcBool(filepath.Join(procSysNetPath, "ipv4/ip_forward"))

	if _, err = os.Stat(appArmorPath); !os.IsNotExist(err) {
		if _, err = os.ReadFile(filepath.Join(appArmorPath, "profiles")); err == nil {
			sysInfo.AppArmor = true
		}
	}

	sysInfo.Seccomp = seccomp.IsEnabled()

	return sysInfo, warnings, nil
}

func readProcBool(path string) bool {
	val, err := os.ReadFile(path)
	if err != nil {
		return false
	}

	return strings.TrimSpace(string(val)) == "1"
}
