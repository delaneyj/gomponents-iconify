package material_symbols

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func Atr(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<path fill="currentColor" d="M6 20q-1.25 0-2.125-.875T3 17q0-1.25.875-2.125T6 14q1.25 0 2.125.875T9 17q0 1.25-.875 2.125T6 20Zm12 0q-1.25 0-2.125-.875T15 17q0-1.25.875-2.125T18 14q1.25 0 2.125.875T21 17q0 1.25-.875 2.125T18 20Zm-6-10q-1.25 0-2.125-.875T9 7q0-1.25.875-2.125T12 4q1.25 0 2.125.875T15 7q0 1.25-.875 2.125T12 10Z"/>`),
		g.Group(children),
	)
}