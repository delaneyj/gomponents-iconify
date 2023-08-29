package icon_park_outline

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func ChartGraph(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 48 48"),
		g.Raw(`<g fill="none" stroke="currentColor" stroke-linejoin="round" stroke-width="4"><path d="M17 6h14v9H17zM6 33h14v9H6zm22 0h14v9H28z"/><path stroke-linecap="round" d="M24 16v8m-11 9v-9h22v9"/></g>`),
		g.Group(children),
	)
}