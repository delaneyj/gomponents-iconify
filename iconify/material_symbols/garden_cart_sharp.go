package material_symbols

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func GardenCartSharp(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<path fill="currentColor" d="M7.825 21H5V9.4L3.175 5H1V3h3.525l1.65 4h16.4L19 14.025q1.275.2 2.138 1.175T22 17.5q0 1.45-1.012 2.475T18.524 21q-1.475 0-2.487-1.025T15.025 17.5q0-.5.125-.925t.35-.825l-3.275-.3l-4.4 5.55ZM7 13.7v5.275l2.975-3.725l-2.425-.225L7 13.7ZM18.5 19q.65 0 1.075-.438T20 17.5q0-.65-.425-1.075T18.5 16q-.625 0-1.063.425T17 17.5q0 .625.438 1.063T18.5 19Z"/>`),
		g.Group(children),
	)
}