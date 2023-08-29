package material_symbols

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func VideoLibraryOutlineSharp(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<path fill="currentColor" d="m11.5 14.5l7-4.5l-7-4.5v9ZM6 18V2h16v16H6Zm2-2h12V4H8v12Zm-6 6V6h2v14h14v2H2ZM8 4v12V4Z"/>`),
		g.Group(children),
	)
}