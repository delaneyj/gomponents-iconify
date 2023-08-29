package icon_park_outline

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func ChartProportion(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 48 48"),
		g.Raw(`<g fill="none" stroke="currentColor" stroke-linejoin="round" stroke-width="4"><path stroke-linecap="round" d="M16.345 5.518a20.029 20.029 0 0 0-6.487 4.34A19.938 19.938 0 0 0 4 24c0 11.046 8.954 20 20 20v0a19.937 19.937 0 0 0 14.142-5.858a20.026 20.026 0 0 0 4.34-6.487"/><path d="M24 24h20c0-11.046-8.954-20-20-20v20Z"/></g>`),
		g.Group(children),
	)
}