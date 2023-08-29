package material_symbols

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func EnergyProgramTimeUsedOutlineSharp(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<path fill="currentColor" d="M10 13Zm-9 9V4h10v2H3v14h14v-6h2v8H1Zm4-4h2v-7H5v7Zm4 0h2V8H9v10Zm4 0h2v-4h-2v4Zm5-6q-.75 0-1.475-.225t-1.35-.65l-1.1 1.025l-1.375-1.375l1.1-1.1q-.4-.6-.6-1.275T13 7q0-2.075 1.463-3.538T18 2h5v5q0 2.075-1.463 3.538T18 12Zm0-2q1.25 0 2.125-.875T21 7V4h-3q-1.25 0-2.125.875T15 7q0 .325.063.625t.187.6l3.3-3.3l1.425 1.4L16.65 9.65q.325.15.663.25T18 10Zm-.125-2.975Z"/>`),
		g.Group(children),
	)
}