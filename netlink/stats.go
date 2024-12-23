package netlink

import "github.com/vishvananda/netlink"

func StatsForLinks(links []netlink.Link) (float64, float64) {
	var received, transmitted float64

	for _, l := range links {
		stats := l.Attrs().Statistics
		if stats != nil {
			received += float64(stats.RxBytes)
			transmitted += float64(stats.TxBytes)
		}
	}

	return received, transmitted
}
