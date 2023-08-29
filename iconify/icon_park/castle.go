package icon_park

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func Castle(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 48 48"),
		g.Raw(`<g fill="none" stroke="#000" stroke-linecap="round" stroke-linejoin="round" stroke-width="4"><path d="M4 11L11 4L18 11H4Z"/><path d="M30 11L37 4L44 11H30Z"/><path fill="#2F88FF" d="M44 44V26H40V17H34V26H28V21L24 17L20 21V26H14V17H8V26H4V44H18V40C18 36.6863 20.6863 34 24 34C27.3137 34 30 36.6863 30 40V44H44Z"/><rect width="8" height="6" x="7" y="11"/><rect width="8" height="6" x="33" y="11"/></g>`),
		g.Group(children),
	)
}