package icon_park_outline

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func TicketOne(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 48 48"),
		g.Raw(`<g fill="none" stroke="currentColor" stroke-linecap="round" stroke-width="4"><path stroke-linejoin="round" d="M4 8h7s1 5 6 5s6-5 6-5h21v32H23s-1-5-6-5s-6 5-6 5H4V8Z"/><path d="M17 19v2m0 6v2m8-8h11m-11 6h11"/></g>`),
		g.Group(children),
	)
}