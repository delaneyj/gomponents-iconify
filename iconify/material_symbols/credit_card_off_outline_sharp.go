package material_symbols

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func CreditCardOffOutlineSharp(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<path fill="currentColor" d="m22 19.15l-2-2V12h-5.15l-4-4H20V6H8.85l-2-2H22v15.15Zm-12.375-6.7Zm4.8-.875ZM9.15 12H4v6h11.15l-6-6Zm11.3 11.3l-3.3-3.3H2V4h2l2 2H4v2h1.15L.65 3.5l1.425-1.425l19.8 19.8L20.45 23.3Z"/>`),
		g.Group(children),
	)
}