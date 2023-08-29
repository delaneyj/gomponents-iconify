package fe

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func LineChart(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<g id="feLineChart0" fill="none" fill-rule="evenodd" stroke="none" stroke-width="1"><g id="feLineChart1" fill="currentColor"><path id="feLineChart2" d="M5.823 14.177a2 2 0 1 1-1-1l2.354-2.354a2 2 0 1 1 3.646 0l2.354 2.354a1.993 1.993 0 0 1 1.646 0l3.354-3.354a2 2 0 1 1 1 1l-3.354 3.354a2 2 0 1 1-3.646 0l-2.354-2.354a1.993 1.993 0 0 1-1.646 0l-2.354 2.354Z"/></g></g>`),
		g.Group(children),
	)
}