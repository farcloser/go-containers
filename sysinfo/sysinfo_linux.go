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
