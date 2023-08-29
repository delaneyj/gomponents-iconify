package icon_park

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func Text(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 48 48"),
		g.Raw(`<g fill="none" stroke-linejoin="round" stroke-width="4"><rect width="36" height="36" x="6" y="6" fill="#2F88FF" stroke="#000" rx="3"/><path stroke="#fff" stroke-linecap="round" d="M16 19V16H32V19"/><path stroke="#fff" stroke-linecap="round" d="M22 34H26"/><path stroke="#fff" stroke-linecap="round" d="M24 18L24 34"/></g>`),
		g.Group(children),
	)
}