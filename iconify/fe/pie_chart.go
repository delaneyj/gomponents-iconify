package fe

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func PieChart(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<g id="fePieChart0" fill="none" fill-rule="evenodd" stroke="none" stroke-width="1"><g id="fePieChart1" fill="currentColor"><path id="fePieChart2" d="M12 22C6.477 22 2 17.523 2 12S6.477 2 12 2s10 4.477 10 10s-4.477 10-10 10Zm0-18a8 8 0 1 0 8 8h-8V4Z"/></g></g>`),
		g.Group(children),
	)
}