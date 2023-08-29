package icon_park_outline

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func BottomBarOne(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 48 48"),
		g.Raw(`<path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="4" d="M13.777 20h4s7 7.28 7 14s-6 10-6 10h-6s-6-3.28-6-10c0-6.72 7-14 7-14Zm2-16c2.084 0 5 1.52 5 6s-3.333 10-3.333 10h-3.333s-3.334-5.52-3.334-10s2.917-6 5-6Zm16 16h4s7 7.28 7 14s-6 10-6 10h-6s-6-3.28-6-10c0-6.72 7-14 7-14Zm2-16c2.084 0 5 1.52 5 6s-3.333 10-3.333 10h-3.333s-3.334-5.52-3.334-10s2.917-6 5-6Z"/>`),
		g.Group(children),
	)
}