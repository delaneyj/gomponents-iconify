package ri

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func PieChartBoxFill(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<path fill="currentColor" d="M3 3h18a1 1 0 0 1 1 1v16a1 1 0 0 1-1 1H3a1 1 0 0 1-1-1V4a1 1 0 0 1 1-1Zm13.9 10H11V7.1a5.002 5.002 0 0 0 1 9.9a5.002 5.002 0 0 0 4.9-4Zm0-2A5.006 5.006 0 0 0 13 7.1V11h3.9Z"/>`),
		g.Group(children),
	)
}