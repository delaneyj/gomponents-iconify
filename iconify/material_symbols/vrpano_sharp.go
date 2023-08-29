package material_symbols

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func VrpanoSharp(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<path fill="currentColor" d="M5.5 15.4q1.5-.2 3.113-.3T12 15q1.725 0 3.363.1t3.137.3L14 10l-2.85 3.4l-2-2.4l-3.65 4.4Zm-3.475 5.1V3.475q1.875.875 4.5 1.45t5.5.575q2.875 0 5.5-.575t4.5-1.45V20.5q-1.875-.875-4.5-1.437t-5.5-.563q-2.875 0-5.5.563t-4.5 1.437Z"/>`),
		g.Group(children),
	)
}