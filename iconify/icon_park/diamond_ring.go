package icon_park

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func DiamondRing(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 48 48"),
		g.Raw(`<g fill="none" stroke="#000" stroke-linecap="round" stroke-linejoin="round" stroke-width="4"><circle cx="25" cy="29" r="15"/><path fill="#2F88FF" d="M18 8L21 4H25.1339H29.0536L32 8L25 14L18 8Z"/></g>`),
		g.Group(children),
	)
}