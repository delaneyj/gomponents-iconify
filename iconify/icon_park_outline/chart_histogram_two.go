package icon_park_outline

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func ChartHistogramTwo(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 48 48"),
		g.Raw(`<g fill="none"><path d="M4 42h40H4Z" clip-rule="evenodd"/><path stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="4" d="M4 42h40"/><path stroke="currentColor" stroke-linejoin="round" stroke-width="4" d="M8 28h6v14H8zm13-10h6v24h-6zM34 6h6v36h-6z"/></g>`),
		g.Group(children),
	)
}