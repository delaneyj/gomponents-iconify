package icon_park_twotone

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func TicketsTwo(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 48 48"),
		g.Raw(`<mask id="ipTTicketsTwo0"><g fill="none"><path stroke="#fff" stroke-linecap="round" stroke-linejoin="round" stroke-width="4" d="M30 30c3 0 5 1.88 5 4h6V4h-6c0 2-2 4-5 4s-5-2-5-4h-6v14"/><path fill="#555" stroke="#fff" stroke-linecap="round" stroke-linejoin="round" stroke-width="4" d="M30 44h-6c0-2.12-2-4-5-4s-5 1.88-5 4H8V14h6c0 2 2 4 5 4s5-2 5-4h6v30Z"/><circle cx="14" cy="24" r="2" fill="#fff"/><circle cx="19" cy="24" r="2" fill="#fff"/><circle cx="24" cy="24" r="2" fill="#fff"/></g></mask><path fill="currentColor" d="M0 0h48v48H0z" mask="url(#ipTTicketsTwo0)"/>`),
		g.Group(children),
	)
}