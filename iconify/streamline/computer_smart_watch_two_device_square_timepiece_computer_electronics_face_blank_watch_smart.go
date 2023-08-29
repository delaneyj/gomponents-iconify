package streamline

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func ComputerSmartWatchTwoDeviceSquareTimepieceComputerElectronicsFaceBlankWatchSmart(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 14 14"),
		g.Raw(`<g fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round"><path d="M9 3.5v-2a1 1 0 0 0-1-1H6a1 1 0 0 0-1 1v2m4 7v2a1 1 0 0 1-1 1H6a1 1 0 0 1-1-1v-2"/><rect width="7" height="7" x="3.5" y="3.5" rx="1"/></g>`),
		g.Group(children),
	)
}