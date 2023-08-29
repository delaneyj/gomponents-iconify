package tdesign

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func ChartAnalytics(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<path fill="currentColor" d="M4 2v18h18v2H2V2h2Zm12 4h6v6h-2V9.423l-1.579 1.575a4780.256 4780.256 0 0 0-2.714 2.708l-.707.707l-4-4l-5 5L4.586 14L11 7.585l4 4l2.009-2.003L18.594 8H16V6Z"/>`),
		g.Group(children),
	)
}