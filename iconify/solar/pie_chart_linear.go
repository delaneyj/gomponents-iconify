package solar

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func PieChartLinear(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<g fill="none" stroke="currentColor" stroke-width="1.5"><path stroke-linecap="round" d="M20 15.552A9.215 9.215 0 0 1 11.21 22A9.21 9.21 0 0 1 2 12.79A9.215 9.215 0 0 1 8.448 4"/><path d="M21.913 9.947a11.352 11.352 0 0 0-7.86-7.86C12.409 1.628 11 3.054 11 4.76v6.694c0 .853.692 1.545 1.545 1.545h6.694c1.707 0 3.133-1.41 2.674-3.053Z"/></g>`),
		g.Group(children),
	)
}