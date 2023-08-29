package material_symbols

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func FullStackedBarChart(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<path fill="currentColor" d="M4 20v-3h4v3H4Zm0-4v-4h4v4H4Zm0-5V4h4v7H4Zm6 9v-7h4v7h-4Zm0-8V8h4v4h-4Zm0-5V4h4v3h-4Zm6 13v-2h4v2h-4Zm0-3v-4h4v4h-4Zm0-5V4h4v8h-4Z"/>`),
		g.Group(children),
	)
}