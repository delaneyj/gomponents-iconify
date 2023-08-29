package material_symbols

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func RoundaboutLeftSharp(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<path fill="currentColor" d="M15 21v-8.05h1q1.65 0 2.825-1.15T20 9q0-1.65-1.175-2.825T16 5q-1.65 0-2.8 1.175T12.05 9v1H5.825l1.6 1.6L6 13L2 9l4-4l1.4 1.4L5.825 8h4.325q.35-2.125 1.988-3.563T16 3q2.5 0 4.25 1.75T22 9q0 2.225-1.438 3.863T17 14.85V21h-2Z"/>`),
		g.Group(children),
	)
}