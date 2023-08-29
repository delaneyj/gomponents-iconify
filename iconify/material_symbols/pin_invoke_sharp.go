package material_symbols

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func PinInvokeSharp(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<path fill="currentColor" d="M2 20V4h12v2H4v12h16v-6h2v8H2Zm8.05-3.625l-1.4-1.425L11.6 12H9.35v-2H15v5.65h-2v-2.225l-2.95 2.95ZM19 10q-1.25 0-2.125-.875T16 7q0-1.25.875-2.125T19 4q1.25 0 2.125.875T22 7q0 1.25-.875 2.125T19 10Z"/>`),
		g.Group(children),
	)
}