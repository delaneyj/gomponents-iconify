package icon_park

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func ViewGridList(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 48 48"),
		g.Raw(`<g fill="none" stroke-linejoin="round" stroke-width="4"><rect width="36" height="36" x="6" y="6" fill="#2F88FF" stroke="#000" rx="3"/><rect width="8" height="8" x="13" y="13" fill="#43CCF8" stroke="#fff"/><rect width="8" height="8" x="13" y="27" fill="#43CCF8" stroke="#fff"/><path stroke="#fff" stroke-linecap="round" d="M27 28L35 28"/><path stroke="#fff" stroke-linecap="round" d="M27 35H35"/><path stroke="#fff" stroke-linecap="round" d="M27 13L35 13"/><path stroke="#fff" stroke-linecap="round" d="M27 20L35 20"/></g>`),
		g.Group(children),
	)
}