package mdi

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func ShapeRectanglePlus(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<path fill="currentColor" d="M19 6h3v2h-3v3h-2V8h-3V6h3V3h2v3m-2 11v-3h2v5H3V6h8v2H5v9h12Z"/>`),
		g.Group(children),
	)
}