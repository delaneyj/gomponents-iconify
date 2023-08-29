package material_symbols

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func IronOutlineSharp(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<path fill="currentColor" d="M2 18v-3q0-1.65 1.175-2.825T6 11h9V9H9v1H7V7h10v7h1V6h4v2h-2v8h-3v2H2Zm2-2h11v-3H6q-.825 0-1.413.588T4 15v1Zm11 0v-3v3Z"/>`),
		g.Group(children),
	)
}