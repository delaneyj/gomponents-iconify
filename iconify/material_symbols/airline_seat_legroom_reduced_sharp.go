package material_symbols

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func AirlineSeatLegroomReducedSharp(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<path fill="currentColor" d="M14 21v-3l1-4H6V3h6v6h6.25l.95 1.275L17 18h3v3h-6Zm-2-4H3V3h2v12h7v2Z"/>`),
		g.Group(children),
	)
}