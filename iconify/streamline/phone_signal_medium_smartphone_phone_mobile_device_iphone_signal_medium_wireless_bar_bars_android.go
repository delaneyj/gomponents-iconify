package streamline

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func PhoneSignalMediumSmartphonePhoneMobileDeviceIphoneSignalMediumWirelessBarBarsAndroid(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 14 14"),
		g.Raw(`<path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" d="M4.75 7.5H.5v6h4.25m0-9h4.5v9h-4.5zm4.5 9h4.25"/>`),
		g.Group(children),
	)
}