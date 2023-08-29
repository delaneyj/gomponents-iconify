package humbleicons

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func Arrows(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<g fill="none" stroke="currentColor" stroke-width="2"><path d="M12 4v6m0 4v6"/><path stroke-linecap="round" stroke-linejoin="round" d="M9.5 6.5L12 4l2.5 2.5m-5 11L12 20l2.5-2.5m-8-8L4 12l2.5 2.5m11-5L20 12l-2.5 2.5M5.5 12h13"/></g>`),
		g.Group(children),
	)
}