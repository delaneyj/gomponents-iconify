package icon_park

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func BringForward(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 48 48"),
		g.Raw(`<g fill="none" stroke="#000" stroke-linejoin="round" stroke-width="4"><rect width="36" height="8" x="6" y="34" fill="#2F88FF"/><rect width="36" height="8" x="6" y="20" fill="#2F88FF"/><path stroke-linecap="round" d="M30 12L24 6L18 12V12"/><path stroke-linecap="round" d="M24 28V34"/><path stroke-linecap="round" d="M24 6V20"/></g>`),
		g.Group(children),
	)
}