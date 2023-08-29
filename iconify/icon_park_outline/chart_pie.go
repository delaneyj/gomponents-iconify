package icon_park_outline

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func ChartPie(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 48 48"),
		g.Raw(`<g fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="4"><path d="M44 24c0 11.046-8.954 20-20 20S4 35.046 4 24S12.954 4 24 4v20h20Z"/><path d="M43.084 18H30V4.916A20.047 20.047 0 0 1 43.084 18Z"/></g>`),
		g.Group(children),
	)
}