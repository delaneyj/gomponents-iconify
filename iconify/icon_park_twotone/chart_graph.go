package icon_park_twotone

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func ChartGraph(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 48 48"),
		g.Raw(`<mask id="ipTChartGraph0"><g fill="none" stroke="#fff" stroke-linejoin="round" stroke-width="4"><path fill="#555" d="M17 6h14v9H17zM6 33h14v9H6zm22 0h14v9H28z"/><path stroke-linecap="round" d="M24 16v8m-11 9v-9h22v9"/></g></mask><path fill="currentColor" d="M0 0h48v48H0z" mask="url(#ipTChartGraph0)"/>`),
		g.Group(children),
	)
}