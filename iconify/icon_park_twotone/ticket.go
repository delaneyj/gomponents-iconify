package icon_park_twotone

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func Ticket(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 48 48"),
		g.Raw(`<mask id="ipTTicket0"><g fill="none" stroke="#fff" stroke-linecap="round" stroke-width="4"><path stroke-linejoin="round" d="M9 16L34 6l4 10"/><path fill="#555" stroke-linejoin="round" d="M4 16h40v6c-3 0-6 2-6 5.5s3 6.5 6 6.5v6H4v-6c3 0 6-2 6-6s-3-6-6-6v-6Z"/><path d="M17 25.385h6m-6 6h14"/></g></mask><path fill="currentColor" d="M0 0h48v48H0z" mask="url(#ipTTicket0)"/>`),
		g.Group(children),
	)
}