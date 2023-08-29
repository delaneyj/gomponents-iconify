package material_symbols

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func SignpostOutline(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<path fill="currentColor" d="M11 22v-4H6l-3-3l3-3h5v-2H4V4h7V2h2v2h5l3 3l-3 3h-5v2h7v6h-7v4h-2ZM6 8h11.175l1-1l-1-1H6v2Zm.825 8H18v-2H6.825l-1 1l1 1ZM6 8V6v2Zm12 8v-2v2Z"/>`),
		g.Group(children),
	)
}