package ri

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func LineChartLine(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<path fill="currentColor" d="M5 3v16h16v2H3V3h2Zm15.293 3.293l1.414 1.414L16 13.414l-3-2.999l-4.293 4.292l-1.414-1.414L13 7.586l3 2.999l4.293-4.292Z"/>`),
		g.Group(children),
	)
}