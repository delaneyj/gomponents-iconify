package ri

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func BarChartTwoLine(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<path fill="currentColor" d="M2 13h6v8H2v-8Zm14-5h6v13h-6V8ZM9 3h6v18H9V3ZM4 15v4h2v-4H4Zm7-10v14h2V5h-2Zm7 5v9h2v-9h-2Z"/>`),
		g.Group(children),
	)
}