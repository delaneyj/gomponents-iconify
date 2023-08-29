package material_symbols

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func QuickReference(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<path fill="currentColor" d="M16.5 20h1v-4h-1v4Zm.5-5q.2 0 .35-.15t.15-.35q0-.2-.15-.35T17 14q-.2 0-.35.15t-.15.35q0 .2.15.35T17 15Zm0 7q-2.075 0-3.538-1.463T12 17q0-2.075 1.463-3.538T17 12q2.075 0 3.538 1.463T22 17q0 2.075-1.463 3.538T17 22ZM12 9h5l-5-5l5 5l-5-5v5ZM5 22q-.825 0-1.413-.588T3 20V4q0-.825.588-1.413T5 2h8l6 6v2.3q-.5-.15-1-.225T17 10q-1.425 0-2.688.538T12.1 12H7v2h3.675q-.225.475-.375.975T10.075 16H7v2h3.075q.175 1.125.7 2.163T12.125 22H5Z"/>`),
		g.Group(children),
	)
}