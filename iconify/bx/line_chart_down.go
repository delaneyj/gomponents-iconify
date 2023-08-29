package bx

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func LineChartDown(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<path fill="currentColor" d="M5 3H3v18h18v-2H5z"/><path fill="currentColor" d="M13 12.586L8.707 8.293L7.293 9.707L13 15.414l3-3l4.293 4.293l1.414-1.414L16 9.586z"/>`),
		g.Group(children),
	)
}