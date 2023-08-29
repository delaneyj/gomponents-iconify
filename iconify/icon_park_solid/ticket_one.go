package icon_park_solid

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func TicketOne(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 48 48"),
		g.Raw(`<mask id="ipSTicketOne0"><g fill="none" stroke-linecap="round" stroke-width="4"><path fill="#fff" stroke="#fff" stroke-linejoin="round" d="M4 8h7s1 5 6 5s6-5 6-5h21v32H23s-1-5-6-5s-6 5-6 5H4V8Z"/><path stroke="#000" d="M17 19v2m0 6v2m8-8h11m-11 6h11"/></g></mask><path fill="currentColor" d="M0 0h48v48H0z" mask="url(#ipSTicketOne0)"/>`),
		g.Group(children),
	)
}