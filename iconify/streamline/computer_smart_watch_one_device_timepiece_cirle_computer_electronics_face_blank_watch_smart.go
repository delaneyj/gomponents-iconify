package streamline

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func ComputerSmartWatchOneDeviceTimepieceCirleComputerElectronicsFaceBlankWatchSmart(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 14 14"),
		g.Raw(`<g fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round"><circle cx="7" cy="7" r="3.5"/><path d="M9 4.13V1.5a1 1 0 0 0-1-1H6a1 1 0 0 0-1 1v2.63m4 5.74v2.63a1 1 0 0 1-1 1H6a1 1 0 0 1-1-1V9.87"/></g>`),
		g.Group(children),
	)
}