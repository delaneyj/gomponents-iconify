package icon_park

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func List(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 48 48"),
		g.Raw(`<g fill="none" stroke="#000" stroke-linecap="round" stroke-linejoin="round" stroke-width="4"><path d="M5 10L8 13L14 7"/><path d="M5 24L8 27L14 21"/><path d="M5 38L8 41L14 35"/><path d="M21 24H43"/><path d="M21 38H43"/><path d="M21 10H43"/></g>`),
		g.Group(children),
	)
}