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

func LinkByName(netInterface string) (Link, error) { //nolint:ireturn
	link, err := netlink.LinkByName(netInterface)
	if err != nil {
		err = errors.Join(ErrLinkNotFound, err)
	}

	return link, err
}

func LinkDel(netInterface string) error {
	link, err := LinkByName(netInterface)
	if err != nil {
		return err
	}

	err = netlink.LinkDel(link)
	if err != nil {
		err = errors.Join(ErrRemoveFail, err)
	}

	return err
}

func GetNetLinks(pid int, interfaces []net.Interface) (nlinks []netlink.Link, err error) {
	var (
		nlink    netlink.Link
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

	for _, v := range interfaces {
		nlink, err = nlHandle.LinkByIndex(v.Index)
		if err != nil {
			err = fmt.Errorf("failed to retrieve the statistics for %s in netns %s: %w", v.Name, nsHandle, err)

			return nlinks, err
		}
		// exclude inactive interface
		if nlink.Attrs().Flags&net.FlagUp != 0 {
			// exclude loopback interface
			if nlink.Attrs().Flags&net.FlagLoopback != 0 || strings.HasPrefix(nlink.Attrs().Name, "lo") {
				continue
			}

			nlinks = append(nlinks, nlink)
		}
	}

	return nlinks, nil
}
