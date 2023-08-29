package iconamoon

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func CursorDuotone(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<g fill="none"><path fill="currentColor" d="M11 21L4 4l17 7l-6.265 2.685a2 2 0 0 0-1.05 1.05L11 21Z" opacity=".16"/><path stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M11 21L4 4l17 7l-6.265 2.685a2 2 0 0 0-1.05 1.05L11 21Z"/></g>`),
		g.Group(children),
	)
}