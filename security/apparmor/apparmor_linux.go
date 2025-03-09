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

package apparmor

import (
	"context"
	"errors"
	"os"
	"os/exec"
	"path/filepath"
	"strings"
	"sync"

	"github.com/containerd/containerd/v2/contrib/apparmor"
	"github.com/containerd/containerd/v2/core/containers"
	"github.com/containerd/containerd/v2/pkg/oci"
	"github.com/moby/sys/userns"
	"github.com/opencontainers/runtime-spec/specs-go"

	"go.farcloser.world/core/filesystem"
)

type mode = string

type Profile struct {
	Name string `json:"name"`
	Mode mode   `json:"mode,omitempty"`
}

const (
	Enforce    = mode("enforce")
	Unconfined = mode("unconfined")
	Complain   = mode("complain")

	kernelPath   = "/sys/kernel/security/apparmor"
	removePath   = "/sys/kernel/security/apparmor/.remove"
	profilesPath = "/sys/kernel/security/apparmor/policy/profiles"
	enabledPath  = "/sys/module/apparmor/parameters/enabled"

	execBinary = "aa-exec"
)

//nolint:gochecknoglobals
var (
	checkAppArmor     sync.Once
	appArmorSupported bool

	checkParamEnabled sync.Once
	paramEnabled      bool
)

// Supported checks whether AppArmnor is supported on the host
// Note this may not be accessible from user namespaces.
func Supported() bool {
	checkAppArmor.Do(func() {
		if _, err := os.Stat(kernelPath); err == nil {
			appArmorSupported = true
		}
	})

	return appArmorSupported
}

// Enabled checks whether AppArmnor is enabled.
func Enabled() bool {
	checkParamEnabled.Do(func() {
		buf, err := os.ReadFile(enabledPath)
		paramEnabled = err == nil && len(buf) > 1 && buf[0] == 'Y'
	})

	return paramEnabled
}

// CanLoadProfile checks if we can load a new profile. This requires root and full access.
func CanLoadProfile() bool {
	// In some rare circumstances, apparmor may be enabled, but the tooling could be missing
	// containerd implementation shells out to aa-parser, so, require it here.
	// See https://github.com/containerd/nerdctl/issues/3945 for details.
	canLoad := true

	pth, err := exec.LookPath("apparmor_parser")
	if err != nil {
		canLoad = false
	}

	if _, err = os.Stat(pth); err != nil {
		canLoad = false
	}

	return !userns.RunningInUserNS() && os.Geteuid() == 0 && Supported() && Enabled() && canLoad
}

// CanApplyProfile checks if we can apply an already loaded profile.
// Note that this does not require access to profilesPath.
func CanApplyProfile(profileName string) bool {
	if !Enabled() {
		return false
	}

	cmd := exec.Command(execBinary, "-p", profileName, "--", "true")
	_, err := cmd.CombinedOutput()

	return err == nil
}

// DumpCurrentProfileAs shows the content of the *current* profile.
func DumpCurrentProfileAs(name string) (string, error) {
	return apparmor.DumpDefaultProfile(name)
}

// LoadDefaultProfileAs loads the content of the *default* profile as `name`.
func LoadDefaultProfileAs(name string) error {
	return apparmor.LoadDefaultProfile(name)
}

// UnloadProfile needs access to /sys/kernel/security/apparmor/.remove .
func UnloadProfile(name string) error {
	// FIXME: not safe
	remover, err := os.OpenFile(
		removePath,
		os.O_RDWR|os.O_TRUNC,
		filesystem.FilePermissionsDefault)
	if err != nil {
		return err
	}

	_, err = remover.WriteString(name)

	return errors.Join(err, remover.Close())
}

// WithProfile returns a SpecOpts that attaches the profile to the spec.
func WithProfile(name string) oci.SpecOpts {
	return func(_ context.Context, _ oci.Client, _ *containers.Container, s *specs.Spec) error {
		s.Process.ApparmorProfile = name

		return nil
	}
}

// Profiles return the list of currently loaded profiles.
//
// Root is not needed, but ability to read /sys/kernel/security/apparmor/policy/profiles is.
// This might not be accessible from user namespaces (because securityfs cannot be mounted in a user namespace).
func Profiles() ([]*Profile, error) {
	entries, err := os.ReadDir(profilesPath)
	if err != nil {
		return nil, err
	}

	res := []*Profile{}

	for _, entry := range entries {
		var bytes []byte

		profile := &Profile{}

		bytes, err = os.ReadFile(filepath.Join(profilesPath, entry.Name(), "name"))
		if err != nil {
			continue
		}

		profile.Name = strings.TrimSpace(string(bytes))

		bytes, err = os.ReadFile(filepath.Join(profilesPath, entry.Name(), "mode"))
		if err == nil {
			profile.Mode = strings.TrimSpace(string(bytes))
		}

		res = append(res, profile)
	}

	return res, nil
}
