package icon_park

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func BookmarkOne(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 48 48"),
		g.Raw(`<g fill="none" stroke="#000" stroke-linejoin="round" stroke-width="4"><path stroke-linecap="round" d="M34 10V4H8V38L14 35"/><path fill="#2F88FF" d="M14 44V10H40V44L27 37.7273L14 44Z"/></g>`),
		g.Group(children),
	)
}