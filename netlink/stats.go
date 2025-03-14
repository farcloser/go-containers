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
