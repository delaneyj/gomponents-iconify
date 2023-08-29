package streamline

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func MoneyGraphAnalyticsBusinessProductGraphDataChartAnalysis(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 14 14"),
		g.Raw(`<g fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round"><path d="M.5.5v13h13"/><path d="M3.5 6.5L6 9l4-6l3.5 2.5"/></g>`),
		g.Group(children),
	)
}