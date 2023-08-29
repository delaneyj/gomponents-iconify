package material_symbols

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func ShoppingBasketOutlineSharp(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<path fill="currentColor" d="M4.025 21L.675 9H6.75l5.225-7.775L17.2 9h6.125l-3.35 12H4.025ZM5.5 19h13l2.2-8H3.3l2.2 8Zm6.5-2q.825 0 1.413-.588T14 15q0-.825-.588-1.413T12 13q-.825 0-1.413.588T10 15q0 .825.588 1.413T12 17ZM9.175 9H14.8l-2.825-4.2l-2.8 4.2ZM12 15Z"/>`),
		g.Group(children),
	)
}