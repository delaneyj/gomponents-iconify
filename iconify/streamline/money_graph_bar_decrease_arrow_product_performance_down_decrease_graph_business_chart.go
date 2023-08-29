package streamline

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func MoneyGraphBarDecreaseArrowProductPerformanceDownDecreaseGraphBusinessChart(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 14 14"),
		g.Raw(`<path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" d="m1.24.5l11.5 5.23m-2.15.81l2.15-.81l-.8-2.16M1.25 6h1.5a.5.5 0 0 1 .5.5v7h0h-2.5h0v-7a.5.5 0 0 1 .5-.5Zm5 1.5h1.5a.5.5 0 0 1 .5.5v5.5h0h-2.5h0V8a.5.5 0 0 1 .5-.5Zm5 1.5h1.5a.5.5 0 0 1 .5.5v4h0h-2.5h0v-4a.5.5 0 0 1 .5-.5Z"/>`),
		g.Group(children),
	)
}