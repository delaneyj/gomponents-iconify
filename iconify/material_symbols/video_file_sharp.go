package material_symbols

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func VideoFileSharp(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<path fill="currentColor" d="M4 22V2h10l6 6v14H4Zm9-13h5l-5-5v5Zm-5 9h6v-2l2 1.05v-4.1L14 14v-2H8v6Z"/>`),
		g.Group(children),
	)
}