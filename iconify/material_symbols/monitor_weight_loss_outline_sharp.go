package material_symbols

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func MonitorWeightLossOutlineSharp(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<path fill="currentColor" d="M15 18h8v2h-8v-2ZM5 19V5v14Zm-2 2V3h18v10h-2V5H5v14h8v2H3Zm9-9q1.25 0 2.125-.875T15 9q0-1.25-.875-2.125T12 6q-1.25 0-2.125.875T9 9q0 1.25.875 2.125T12 12Zm-2-2.5v-1h1v1h-1Zm1.5 0v-1h1v1h-1Zm1.5 0v-1h1v1h-1Z"/>`),
		g.Group(children),
	)
}