package material_symbols

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func ContactsSharp(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<path fill="currentColor" d="M4 23v-2h16v2H4ZM4 3V1h16v2H4Zm8 10q1.25 0 2.125-.875T15 10q0-1.25-.875-2.125T12 7q-1.25 0-2.125.875T9 10q0 1.25.875 2.125T12 13ZM2 20V4h20v16H2Zm3.75-2h12.5q-1.125-1.4-2.725-2.2T12 15q-1.925 0-3.525.8T5.75 18Z"/>`),
		g.Group(children),
	)
}