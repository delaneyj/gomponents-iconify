package icon_park

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func QuadrangularPyramid(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 48 48"),
		g.Raw(`<g fill="none" stroke="#000" stroke-linecap="round" stroke-linejoin="round" stroke-width="4"><path d="M44 24L24 4L4 24L24 44L44 24Z"/><path d="M4 24H44"/><path d="M24 44L18 24L24 4"/><path d="M24 44L30 24L24 4"/></g>`),
		g.Group(children),
	)
}