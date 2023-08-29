package material_symbols

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func ShoppingBagSharp(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<path fill="currentColor" d="M4 22V6h4q0-1.65 1.175-2.825T12 2q1.65 0 2.825 1.175T16 6h4v16H4Zm6-16h4q0-.825-.588-1.413T12 4q-.825 0-1.413.588T10 6Zm-2 5h2V8H8v3Zm6 0h2V8h-2v3Z"/>`),
		g.Group(children),
	)
}