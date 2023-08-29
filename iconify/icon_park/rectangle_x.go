package icon_park

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func RectangleX(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 48 48"),
		g.Raw(`<g fill="none" stroke="#000" stroke-linejoin="round" stroke-width="4"><path fill="#2F88FF" d="M24 4L4 24L24 44L44 24L24 4Z"/><path stroke-linecap="round" d="M44 4L4 44"/><path stroke-linecap="round" d="M4 4L44 44"/></g>`),
		g.Group(children),
	)
}