package material_symbols

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func SportsMotorsportsSharp(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<path fill="currentColor" d="M2.075 19.45v-.325q0-1.175.038-2.062t.112-1.588h7.85q1.725 0 2.938-1.225t1.212-3q0-1.275-.712-2.313T11.6 7.4L8.25 6.075q1.35-.75 2.838-1.137t3.087-.388q3.2 0 5.475 2.163T21.925 12q0 3.125-2.163 5.288t-5.287 2.162h-12.4Zm.475-5.95q.5-1.775 1.475-3.337T6.35 7.4l4.5 1.825q.65.275 1.025.825t.375 1.2q0 .95-.637 1.6t-1.588.65H2.55Z"/>`),
		g.Group(children),
	)
}