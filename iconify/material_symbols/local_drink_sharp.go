package material_symbols

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func LocalDrinkSharp(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<path fill="currentColor" d="M5.2 22L3 2h18l-2.2 20H5.2Zm.45-14h12.7l.4-4H5.25l.4 4ZM12 19q1.2 0 2.025-.825t.825-2.025q0-1.025-.663-2.225T12 11q-1.525 1.725-2.188 2.925T9.15 16.15q0 1.2.825 2.025T12 19Z"/>`),
		g.Group(children),
	)
}