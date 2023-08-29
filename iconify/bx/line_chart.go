package bx

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func LineChart(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<path fill="currentColor" d="M3 3v17a1 1 0 0 0 1 1h17v-2H5V3H3z"/><path fill="currentColor" d="M15.293 14.707a.999.999 0 0 0 1.414 0l5-5l-1.414-1.414L16 12.586l-2.293-2.293a.999.999 0 0 0-1.414 0l-5 5l1.414 1.414L13 12.414l2.293 2.293z"/>`),
		g.Group(children),
	)
}