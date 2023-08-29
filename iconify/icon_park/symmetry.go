package icon_park

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func Symmetry(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 48 48"),
		g.Raw(`<g fill="none" stroke="#000" stroke-linejoin="round" stroke-width="4"><path d="M4 15L24 9V39L4 33V15Z"/><path d="M24 9L44 15V33L24 39V9Z"/><path stroke-linecap="round" d="M24 4V44"/></g>`),
		g.Group(children),
	)
}