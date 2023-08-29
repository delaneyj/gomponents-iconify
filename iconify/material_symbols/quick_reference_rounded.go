package material_symbols

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func QuickReferenceRounded(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<path fill="currentColor" d="M17 20q.2 0 .35-.15t.15-.35v-3q0-.2-.15-.35T17 16q-.2 0-.35.15t-.15.35v3q0 .2.15.35T17 20Zm0-5q.2 0 .35-.15t.15-.35q0-.2-.15-.35T17 14q-.2 0-.35.15t-.15.35q0 .2.15.35T17 15Zm0 7q-2.075 0-3.538-1.463T12 17q0-2.075 1.463-3.538T17 12q2.075 0 3.538 1.463T22 17q0 2.075-1.463 3.538T17 22ZM13 9h4l-5-5l5 5l-5-5v4q0 .425.288.713T13 9ZM5 22q-.825 0-1.413-.588T3 20V4q0-.825.588-1.413T5 2h7.175q.4 0 .763.15t.637.425l4.85 4.85q.275.275.425.638t.15.762V10.3q-.5-.15-1-.225T17 10q-1.425 0-2.688.537T12.1 12H8q-.425 0-.713.288T7 13q0 .425.288.713T8 14h2.675q-.225.475-.375.975T10.075 16H8q-.425 0-.713.288T7 17q0 .425.288.713T8 18h2.075q.175 1.125.7 2.163T12.125 22H5Z"/>`),
		g.Group(children),
	)
}