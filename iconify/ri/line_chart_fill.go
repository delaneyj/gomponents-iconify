package ri

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func LineChartFill(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<path fill="currentColor" d="M5 3v16h16v2H3V3h2Zm14.94 2.94l2.12 2.12L16 14.122l-3-3l-3.94 3.94l-2.12-2.122L13 6.88l3 3l3.94-3.94Z"/>`),
		g.Group(children),
	)
}