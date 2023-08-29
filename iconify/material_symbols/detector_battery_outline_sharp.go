package material_symbols

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func DetectorBatteryOutlineSharp(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<path fill="currentColor" d="M9 19h8v-3H9v3Zm-4 2v-7h14v2h2v3h-2v2H5ZM5 5v1h14V5H5Zm3.1 3l.3 1h7.2l.3-1H8.1Zm-1.15 3L6 8H3V3h18v5h-3l-1.15 3h-9.9ZM5 5v1v-1Z"/>`),
		g.Group(children),
	)
}