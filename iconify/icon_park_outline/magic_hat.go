package icon_park_outline

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func MagicHat(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 48 48"),
		g.Raw(`<g fill="none" stroke="currentColor" stroke-width="4"><path stroke-linecap="round" stroke-linejoin="round" d="M10 10c.5 3 1 4.5 1.5 8c.4 2.8.5 7.167.5 9c-2.167 1-7 3-7 7s5 9 19 9s19-5 19-9s-7-7-7-7s0-5.5.5-9s1-5 1.5-8"/><path stroke-linecap="round" stroke-linejoin="round" d="M36 27c0 4-1 8-12.5 8"/><ellipse cx="24" cy="10" rx="14" ry="5"/></g>`),
		g.Group(children),
	)
}