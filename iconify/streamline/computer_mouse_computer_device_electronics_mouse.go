package streamline

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func ComputerMouseComputerDeviceElectronicsMouse(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 14 14"),
		g.Raw(`<g fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round"><rect width="7" height="10.5" x="1" y="3" rx="3.5"/><path d="M1 7h7M4.5 7V2.75A2.25 2.25 0 0 1 6.75.5h0A2.25 2.25 0 0 1 9 2.75V3a2 2 0 0 0 2 2h0a2 2 0 0 0 2-2V1.5"/></g>`),
		g.Group(children),
	)
}