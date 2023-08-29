package material_symbols

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func AssignmentTurnedInOutlineSharp(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<path fill="currentColor" d="M10.6 16.05L17.65 9l-1.4-1.4l-5.65 5.65l-2.85-2.85l-1.4 1.4l4.25 4.25ZM3 21V3h6.2q.325-.9 1.088-1.45T12 1q.95 0 1.713.55T14.8 3H21v18H3Zm2-2h14V5H5v14Zm7-14.75q.325 0 .537-.213t.213-.537q0-.325-.213-.537T12 2.75q-.325 0-.537.213t-.213.537q0 .325.213.537T12 4.25ZM5 19V5v14Z"/>`),
		g.Group(children),
	)
}