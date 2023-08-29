package majesticons

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func TicketTextLine(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M17 19H4a1 1 0 0 1-1-1V6a1 1 0 0 1 1-1h13m0 14h3a1 1 0 0 0 1-1V6a1 1 0 0 0-1-1h-3m0 14v-1m0-13v1m0 3.001V9m0 3.001V12m0 3.001V15M7 12h6m-6 3h3"/>`),
		g.Group(children),
	)
}