package icon_park

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func Box(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 48 48"),
		g.Raw(`<g fill="none" stroke-linejoin="round" stroke-width="4"><rect width="36" height="30" x="6" y="12" fill="#2F88FF" stroke="#000" rx="2"/><path stroke="#fff" stroke-linecap="round" d="M17.9497 24.0083L29.9497 24.0083"/><path stroke="#000" stroke-linecap="round" d="M6 13L13 5H35L42 13"/></g>`),
		g.Group(children),
	)
}