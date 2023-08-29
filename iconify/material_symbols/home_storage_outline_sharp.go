package material_symbols

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func HomeStorageOutlineSharp(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<path fill="currentColor" d="M5 21L3 9h18l-2 12H5Zm1.675-2h10.65l1.275-8H5.4l1.275 8ZM9 15h6v-2H9v2ZM5 8V6h14v2H5Zm2-3V3h10v2H7Zm-.325 14h10.65h-10.65Z"/>`),
		g.Group(children),
	)
}