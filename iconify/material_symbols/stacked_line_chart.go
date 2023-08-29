package material_symbols

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func StackedLineChart(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<path fill="currentColor" d="M3.5 21.5L2 20l7.5-7.5l4 4l7.1-8L22 9.9l-8.5 9.6l-4-4l-6 6Zm0-6L2 14l7.5-7.5l4 4l7.1-8L22 3.9l-8.5 9.6l-4-4l-6 6Z"/>`),
		g.Group(children),
	)
}