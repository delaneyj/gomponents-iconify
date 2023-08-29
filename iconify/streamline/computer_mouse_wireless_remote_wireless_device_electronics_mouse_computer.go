package streamline

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func ComputerMouseWirelessRemoteWirelessDeviceElectronicsMouseComputer(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 14 14"),
		g.Raw(`<g fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round"><rect width="9" height="13" x="2.5" y=".5" rx="4.5"/><path d="M7 3.5V5"/></g>`),
		g.Group(children),
	)
}