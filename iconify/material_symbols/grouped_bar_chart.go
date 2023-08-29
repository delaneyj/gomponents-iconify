package material_symbols

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func GroupedBarChart(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<path fill="currentColor" d="M4 20V8h4v12H4Zm5 0v-7h4v7H9Zm7 0V4h4v16h-4Z"/>`),
		g.Group(children),
	)
}