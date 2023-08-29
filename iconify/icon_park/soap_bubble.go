package icon_park

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func SoapBubble(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 48 48"),
		g.Raw(`<g fill="none" stroke-linecap="round" stroke-width="4"><ellipse cx="22" cy="30" fill="#2F88FF" stroke="#000" stroke-linejoin="round" rx="16" ry="14"/><path stroke="#fff" d="M26 24C27.3333 24.1667 30 25 31 29"/><circle cx="8" cy="8" r="4" fill="#2F88FF" stroke="#000" stroke-linejoin="round"/><circle cx="41" cy="13" r="3" fill="#2F88FF" stroke="#000" stroke-linejoin="round"/></g>`),
		g.Group(children),
	)
}