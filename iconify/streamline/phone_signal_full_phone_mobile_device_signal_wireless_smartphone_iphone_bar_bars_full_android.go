package streamline

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func PhoneSignalFullPhoneMobileDeviceSignalWirelessSmartphoneIphoneBarBarsFullAndroid(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 14 14"),
		g.Raw(`<path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" d="M4.75 7.5H.5v6h4.25m4.5-9h-4.5v9h4.5m0-13h4.25v13H9.25z"/>`),
		g.Group(children),
	)
}