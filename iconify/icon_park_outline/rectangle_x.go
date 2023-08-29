package icon_park_outline

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func RectangleX(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 48 48"),
		g.Raw(`<g fill="none" stroke="currentColor" stroke-linejoin="round" stroke-width="4"><path d="M24 4L4 24l20 20l20-20L24 4Z"/><path stroke-linecap="round" d="M44 4L4 44M4 4l40 40"/></g>`),
		g.Group(children),
	)
}