package icon_park_outline

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func TwoTriangles(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 48 48"),
		g.Raw(`<g fill="none"><path d="m24 4l17.32 30H6.68L24 4Z"/><path d="m24 44l17.32-30H6.68L24 44Z"/><path stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="4" d="m24 4l17.32 30H6.68L24 4Z"/><path stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="4" d="m24 44l17.32-30H6.68L24 44Z"/></g>`),
		g.Group(children),
	)
}