package material_symbols

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func SaveAsSharp(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<path fill="currentColor" d="M3 21V3h14l4 4v5.4L12.4 21H3Zm9-3q1.25 0 2.125-.875T15 15q0-1.25-.875-2.125T12 12q-1.25 0-2.125.875T9 15q0 1.25.875 2.125T12 18Zm-6-8h9V6H6v4Zm9 13v-1.775l5-4.975l1.75 1.775L16.775 23H15Zm7.4-5.65l-1.775-1.75l1.2-1.225L23.6 16.15l-1.2 1.2Z"/>`),
		g.Group(children),
	)
}