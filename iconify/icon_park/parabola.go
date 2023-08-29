package icon_park

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func Parabola(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 48 48"),
		g.Raw(`<g fill="none" stroke="#000" stroke-linecap="round" stroke-linejoin="round" stroke-width="4"><path d="M4 25H6C12 25 16 9 24 9C32 9 36 25 42 25H44"/><path d="M4 33V39"/><path d="M24 33V39"/><path d="M44 33V39"/><path d="M4 36H44"/></g>`),
		g.Group(children),
	)
}