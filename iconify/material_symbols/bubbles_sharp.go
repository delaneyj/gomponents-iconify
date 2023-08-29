package material_symbols

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func BubblesSharp(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<path fill="currentColor" d="M10.05 7.65L13 10.575V8.35h2V14H9.35v-2h2.25L8.65 9.05l1.4-1.4ZM19 20q-1.25 0-2.125-.875T16 17q0-1.25.875-2.125T19 14q1.25 0 2.125.875T22 17q0 1.25-.875 2.125T19 20ZM2 20V4h20v8h-2V6H4v12h10v2H2Z"/>`),
		g.Group(children),
	)
}