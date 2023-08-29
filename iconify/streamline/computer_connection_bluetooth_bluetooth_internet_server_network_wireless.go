package streamline

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func ComputerConnectionBluetoothBluetoothInternetServerNetworkWireless(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 14 14"),
		g.Raw(`<path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" d="m3.25 9.5l7.5-5.5l-4-3.5v13l4-3.5l-7.5-5.5"/>`),
		g.Group(children),
	)
}