package icon_park

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func FenceOne(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 48 48"),
		g.Raw(`<g fill="none" stroke="#000" stroke-linecap="round" stroke-linejoin="round" stroke-width="4"><path d="M8 4V44"/><path d="M44 12L4 12"/><path d="M40 44L8 12"/><path d="M8 44L40 12"/><path d="M27 44L8 25"/><path d="M40 31L21 12"/><path d="M8 31L26 12"/><path d="M21 44L40 25"/><path d="M44 44L4 44"/><path d="M40 4V44"/></g>`),
		g.Group(children),
	)
}