package icon_park

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func AdProduct(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 48 48"),
		g.Raw(`<g fill="none" stroke="#000" stroke-linejoin="round" stroke-width="4"><path d="M44 14L24 4L4 14V34L24 44L44 34V14Z"/><path stroke-linecap="round" d="M4 14L24 24"/><path stroke-linecap="round" d="M24 44V24"/><path stroke-linecap="round" d="M44 14L24 24"/><path stroke-linecap="round" d="M34 9L14 19"/></g>`),
		g.Group(children),
	)
}