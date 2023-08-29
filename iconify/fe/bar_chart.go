package fe

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func BarChart(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<g id="feBarChart0" fill="none" fill-rule="evenodd" stroke="none" stroke-width="1"><g id="feBarChart1" fill="currentColor"><path id="feBarChart2" d="M5 19h15a1 1 0 0 1 0 2H4a1 1 0 0 1-1-1v-8a1 1 0 0 1 2 0v7Zm5-4a1 1 0 0 1-2 0V6a1 1 0 1 1 2 0v9Zm2 0V8a1 1 0 0 1 2 0v7a1 1 0 0 1-2 0Zm4-11a1 1 0 0 1 2 0v11a1 1 0 0 1-2 0V4Z"/></g></g>`),
		g.Group(children),
	)
}