package humbleicons

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func ZoomIn(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<g fill="none" stroke="currentColor" stroke-width="2"><path d="M16 10a6 6 0 1 1-12 0a6 6 0 0 1 12 0Z"/><path stroke-linecap="round" d="m20 20l-5-5m-7.5-5H10m0 0h2.5M10 10v2.5m0-2.5V7.5"/></g>`),
		g.Group(children),
	)
}