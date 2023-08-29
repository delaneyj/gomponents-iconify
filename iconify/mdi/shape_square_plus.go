package mdi

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func ShapeSquarePlus(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<path fill="currentColor" d="M19 5h3v2h-3v3h-2V7h-3V5h3V2h2v3m-2 14v-6h2v8H3V5h8v2H5v12h12Z"/>`),
		g.Group(children),
	)
}