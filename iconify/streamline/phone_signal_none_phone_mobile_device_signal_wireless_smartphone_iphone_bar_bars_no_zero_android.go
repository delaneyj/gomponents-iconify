package streamline

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func PhoneSignalNonePhoneMobileDeviceSignalWirelessSmartphoneIphoneBarBarsNoZeroAndroid(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 14 14"),
		g.Raw(`<path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" d="M.5 13.5H3m8 0h2.5m-8 0h3"/>`),
		g.Group(children),
	)
}