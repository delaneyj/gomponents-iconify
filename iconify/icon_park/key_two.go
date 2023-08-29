package icon_park

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func KeyTwo(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 48 48"),
		g.Raw(`<g fill="none" stroke="#000" stroke-width="4"><circle cx="15" cy="33" r="8" fill="#2F88FF"/><path stroke-linecap="round" stroke-linejoin="round" d="M29 16L36 22"/><path stroke-linecap="round" stroke-linejoin="round" d="M20 26L36 8L43 14"/></g>`),
		g.Group(children),
	)
}