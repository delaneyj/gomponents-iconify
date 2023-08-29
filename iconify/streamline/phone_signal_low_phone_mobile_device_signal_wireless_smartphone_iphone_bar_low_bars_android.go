package streamline

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func PhoneSignalLowPhoneMobileDeviceSignalWirelessSmartphoneIphoneBarLowBarsAndroid(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 14 14"),
		g.Raw(`<path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" d="M.5 7.5h4.25v6H.5zm10.5 6h2.5m-8.75 0H8.5"/>`),
		g.Group(children),
	)
}