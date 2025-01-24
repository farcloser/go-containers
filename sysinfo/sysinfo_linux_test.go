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

/*
   Portions from
	https://github.com/moby/moby/blob/cff4f20c44a3a7c882ed73934dec6a77246c6323/pkg/sysinfo/sysinfo_linux_test.go
   Copyright (C) Docker/Moby authors.
   Licensed under the Apache License, Version 2.0
   NOTICE: https://github.com/moby/moby/blob/cff4f20c44a3a7c882ed73934dec6a77246c6323/NOTICE
*/

package sysinfo_test // import "github.com/docker/docker/pkg/sysinfo"

import (
	"os"
	"testing"

	"golang.org/x/sys/unix"
	"gotest.tools/v3/assert"

	"go.farcloser.world/containers/sysinfo"
)

func TestNew(t *testing.T) {
	t.Parallel()

	sysInfo, _, _ := sysinfo.New("/")
	assert.Assert(t, sysInfo != nil)
	checkSysInfo(t, sysInfo)
}

func checkSysInfo(t *testing.T, sysInfo *sysinfo.SysInfo) {
	t.Helper()

	// Check if Seccomp is supported, via CONFIG_SECCOMP.then sysInfo.Seccomp must be TRUE , else FALSE
	if err := unix.Prctl(unix.PR_GET_SECCOMP, 0, 0, 0, 0); err != unix.EINVAL { //nolint:err113
		// Make sure the kernel has CONFIG_SECCOMP_FILTER.
		if err := unix.Prctl(unix.PR_SET_SECCOMP, unix.SECCOMP_MODE_FILTER, 0, 0, 0); err != unix.EINVAL { //nolint:err113
			assert.Assert(t, sysInfo.Seccomp)
		}
	} else {
		assert.Assert(t, !sysInfo.Seccomp)
	}
}

func TestNewAppArmorEnabled(t *testing.T) {
	t.Parallel()

	// Check if AppArmor is supported. then it must be TRUE , else FALSE
	if _, err := os.Stat("/sys/kernel/security/apparmor"); err != nil {
		t.Skip("AppArmor Must be Enabled")
	}

	// FIXME: rootless is not allowed to read the profile
	if os.Geteuid() != 0 {
		t.Skip("test skipped for rootless")
	}

	sysInfo, _, _ := sysinfo.New("/")
	assert.Assert(t, sysInfo.AppArmor)
}

func TestNewAppArmorDisabled(t *testing.T) {
	t.Parallel()

	// Check if AppArmor is supported. then it must be TRUE , else FALSE
	if _, err := os.Stat("/sys/kernel/security/apparmor"); !os.IsNotExist(err) {
		t.Skip("AppArmor Must be Disabled")
	}

	sysInfo, _, _ := sysinfo.New("/")
	assert.Assert(t, !sysInfo.AppArmor)
}

func TestNewCgroupNamespacesEnabled(t *testing.T) {
	t.Parallel()

	// If cgroup namespaces are supported in the kernel, then sysInfo.CgroupNamespaces should be TRUE
	if _, err := os.Stat("/proc/self/ns/cgroup"); err != nil {
		t.Skip("cgroup namespaces must be enabled")
	}

	sysInfo, _, _ := sysinfo.New("/")
	assert.Assert(t, sysInfo.CgroupNamespaces)
}

func TestNewCgroupNamespacesDisabled(t *testing.T) {
	t.Parallel()

	// If cgroup namespaces are *not* supported in the kernel, then sysInfo.CgroupNamespaces should be FALSE
	if _, err := os.Stat("/proc/self/ns/cgroup"); !os.IsNotExist(err) {
		t.Skip("cgroup namespaces must be disabled")
	}

	sysInfo, _, _ := sysinfo.New("/")
	assert.Assert(t, !sysInfo.CgroupNamespaces)
}

func TestNumCPU(t *testing.T) {
	t.Parallel()

	cpuNumbers := sysinfo.NumCPU()
	if cpuNumbers <= 0 {
		t.Fatal("CPU returned must be greater than zero")
	}
}
