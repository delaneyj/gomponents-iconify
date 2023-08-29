package iconamoon

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func TicketFill(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<path fill="currentColor" fill-rule="evenodd" d="M5 4a3 3 0 0 0-3 3v3a1 1 0 0 0 1 1a1 1 0 0 1 0 2a1 1 0 0 0-1 1v3a3 3 0 0 0 3 3h14a3 3 0 0 0 3-3v-3a1 1 0 0 0-1-1a1 1 0 0 1 0-2a1 1 0 0 0 1-1V7a3 3 0 0 0-3-3H5ZM4 7a1 1 0 0 1 1-1h4v12H5a1 1 0 0 1-1-1v-2.171a2.99 2.99 0 0 0 1.121-.708l-.692-.692l.692.692A3 3 0 0 0 4 9.171V7Z" clip-rule="evenodd"/>`),
		g.Group(children),
	)
}