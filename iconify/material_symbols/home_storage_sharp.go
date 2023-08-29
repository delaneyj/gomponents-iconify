package material_symbols

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func HomeStorageSharp(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<path fill="currentColor" d="M5 21L3 9h18l-2 12H5Zm4-6h6v-2H9v2ZM5 8V6h14v2H5Zm2-3V3h10v2H7Z"/>`),
		g.Group(children),
	)
}