package ri

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func PieChartBoxLine(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<path fill="currentColor" d="M3 3h18a1 1 0 0 1 1 1v16a1 1 0 0 1-1 1H3a1 1 0 0 1-1-1V4a1 1 0 0 1 1-1Zm1 2v14h16V5H4Zm12.9 8A5.002 5.002 0 0 1 7 12a5.002 5.002 0 0 1 4-4.9V13h5.9Zm0-2H13V7.1a5.006 5.006 0 0 1 3.9 3.9Z"/>`),
		g.Group(children),
	)
}