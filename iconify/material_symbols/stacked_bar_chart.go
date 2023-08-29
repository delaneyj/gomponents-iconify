package material_symbols

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func StackedBarChart(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<path fill="currentColor" d="M4 20V9h4v11H4ZM4 8V4h4v4H4Zm6 12v-8h4v8h-4Zm0-9V7h4v4h-4Zm6 9v-5h4v5h-4Zm0-6v-4h4v4h-4Z"/>`),
		g.Group(children),
	)
}