package icon_park

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func ViewGridDetail(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 48 48"),
		g.Raw(`<g fill="none" stroke-linejoin="round" stroke-width="4"><rect width="36" height="36" x="6" y="6" fill="#2F88FF" stroke="#000" rx="3"/><rect width="8" height="8" x="13" y="13" fill="#43CCF8" stroke="#fff"/><path stroke="#fff" stroke-linecap="round" d="M27 13L35 13"/><path stroke="#fff" stroke-linecap="round" d="M27 20L35 20"/><path stroke="#fff" stroke-linecap="round" d="M13 28L35 28"/><path stroke="#fff" stroke-linecap="round" d="M13 35H35"/></g>`),
		g.Group(children),
	)
}