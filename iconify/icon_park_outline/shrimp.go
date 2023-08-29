package icon_park_outline

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func Shrimp(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 48 48"),
		g.Raw(`<g fill="none"><path stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="4" d="M20.623 11H39s0 5-3 8s-6 4-7 4s-10.5 1-13 3s-3.999 5.5 0 9s11 5 16 3s8-8 8-8l4 10s-8 3-15 4s-14 0-20-6s-7.001-13.5-2-20c4.166-5.416 10.414-6.666 12.382-6.933A9.202 9.202 0 0 1 20.623 11Z"/><path stroke="currentColor" stroke-width="4" d="m26 23l-6-12"/><path stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="4" d="M18 4h14a7 7 0 0 1 7 7v0"/><circle cx="30" cy="16" r="2" fill="currentColor"/></g>`),
		g.Group(children),
	)
}