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

package netlink

import (
	"errors"
	"fmt"
	"net"
	"strings"

	"github.com/vishvananda/netlink"
	"github.com/vishvananda/netns"
)

var (
	ErrLinkNotFound = errors.New("link not found")
	ErrRemoveFail   = errors.New("unable to remove network interface")
)

type Link = netlink.Link

func LinkDel(netInterface string) error { //nolint:ireturn,nolintlint // note this is probably a bug in ireturn
	link, err := netlink.LinkByName(netInterface)
	if err != nil {
		return err
	}

	err = netlink.LinkDel(link)
	if err != nil {
		err = errors.Join(ErrRemoveFail, err)
	}

	return err
}

func GetNetNsLinks(pid int) (nlinks []netlink.Link, err error) {
	var (
		nlHandle *netlink.Handle
		nsHandle netns.NsHandle
	)

	nsHandle, err = netns.GetFromPid(pid)
	if err != nil {
		err = fmt.Errorf("failed to retrieve the statistics in netns %s: %w", nsHandle, err)

		return nil, err
	}

	defer func() {
		err = errors.Join(nsHandle.Close(), err)
	}()

	nlHandle, err = netlink.NewHandleAt(nsHandle)
	if err != nil {
		err = fmt.Errorf("failed to retrieve the statistics in netns %s: %w", nsHandle, err)

		return nil, err
	}

	defer nlHandle.Close()

	candidates, err := nlHandle.LinkList()
	if err != nil {
		return nil, err
	}

	for _, nlink := range candidates {
		// exclude down and loopback interfaces
		if nlink.Attrs().Flags&net.FlagUp != 0 {
			if nlink.Attrs().Flags&net.FlagLoopback != 0 || strings.HasPrefix(nlink.Attrs().Name, "lo") {
				continue
			}

			nlinks = append(nlinks, nlink)
		}
	}

	return nlinks, nil
}
