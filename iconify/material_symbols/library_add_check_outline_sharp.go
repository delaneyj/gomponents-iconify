package material_symbols

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func LibraryAddCheckOutlineSharp(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<path fill="currentColor" d="m12.7 14.05l5.65-5.65l-1.4-1.45l-4.25 4.25l-2.15-2.1l-1.4 1.4l3.55 3.55ZM6 18V2h16v16H6Zm2-2h12V4H8v12Zm-6 6V6h2v14h14v2H2ZM8 4v12V4Z"/>`),
		g.Group(children),
	)
}