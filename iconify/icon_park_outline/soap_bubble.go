package icon_park_outline

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func SoapBubble(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 48 48"),
		g.Raw(`<g fill="none" stroke="currentColor" stroke-linecap="round" stroke-width="4"><ellipse cx="22" cy="30" stroke-linejoin="round" rx="16" ry="14"/><path d="M26 24c1.333.167 4 1 5 5"/><circle cx="8" cy="8" r="4" stroke-linejoin="round"/><circle cx="41" cy="13" r="3" stroke-linejoin="round"/></g>`),
		g.Group(children),
	)
}