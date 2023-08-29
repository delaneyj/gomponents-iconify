package material_symbols

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func AirlineSeatLegroomNormalSharp(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<path fill="currentColor" d="M16 21v-7H6V3h6v6h7v9h3v3h-6Zm-2-4H3V3h2v12h9v2Z"/>`),
		g.Group(children),
	)
}