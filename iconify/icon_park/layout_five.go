package icon_park

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func LayoutFive(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 48 48"),
		g.Raw(`<g fill="none" stroke-linejoin="round" stroke-width="4"><rect width="36" height="36" x="6" y="6" fill="#2F88FF" stroke="#000" rx="3"/><path stroke="#fff" stroke-linecap="round" d="M6 16H42"/><path stroke="#000" stroke-linecap="round" d="M6 13V19"/><path stroke="#000" stroke-linecap="round" d="M42 13V19"/><path stroke="#fff" stroke-linecap="round" d="M17 30L42 30"/><path stroke="#fff" stroke-linecap="round" d="M17 16V42"/><path stroke="#000" stroke-linecap="round" d="M14 42H20"/><path stroke="#000" stroke-linecap="round" d="M42 27V33"/></g>`),
		g.Group(children),
	)
}