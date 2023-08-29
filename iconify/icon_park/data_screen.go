package icon_park

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func DataScreen(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 48 48"),
		g.Raw(`<g fill="none" stroke-linecap="round" stroke-linejoin="round" stroke-width="4"><rect width="40" height="30" x="4" y="6" fill="#2F88FF" stroke="#000" rx="3"/><path stroke="#000" d="M24 36V43"/><path stroke="#fff" d="M32 14L16 28"/><path stroke="#000" d="M10 43H38"/><circle cx="15" cy="17" r="3" fill="#43CCF8" stroke="#fff"/><circle cx="33" cy="25" r="3" fill="#43CCF8" stroke="#fff"/></g>`),
		g.Group(children),
	)
}