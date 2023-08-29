package icon_park

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func AddSubtract(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 48 48"),
		g.Raw(`<g fill="none" stroke-linecap="round" stroke-linejoin="round" stroke-width="4"><rect width="36" height="36" x="6" y="6" fill="#2F88FF" stroke="#000" rx="3"/><path stroke="#fff" d="M27 31H35"/><path stroke="#fff" d="M17 13V21"/><path stroke="#fff" d="M21 17H13"/><path stroke="#fff" d="M34 14L14 34"/></g>`),
		g.Group(children),
	)
}