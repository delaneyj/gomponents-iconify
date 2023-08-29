package fe

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func Ticket(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<g id="feTicket0" fill="none" fill-rule="evenodd" stroke="none" stroke-width="1"><g id="feTicket1" fill="currentColor"><path id="feTicket2" d="M2 10V8a2 2 0 0 1 2-2h16a2 2 0 0 1 2 2v2a2 2 0 1 0 0 4v2a2 2 0 0 1-2 2H4a2 2 0 0 1-2-2v-2a2 2 0 1 0 0-4Zm3-2v8h14V8H5Zm2 2h10v4H7v-4Z"/></g></g>`),
		g.Group(children),
	)
}