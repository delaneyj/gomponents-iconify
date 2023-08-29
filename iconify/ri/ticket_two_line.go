package ri

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func TicketTwoLine(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<path fill="currentColor" d="M21.005 3a1 1 0 0 1 1 1v5.5a2.5 2.5 0 0 0 0 5V20a1 1 0 0 1-1 1h-18a1 1 0 0 1-1-1v-5.5a2.5 2.5 0 0 0 0-5V4a1 1 0 0 1 1-1h18Zm-1 2h-16v2.968l.156.081A4.5 4.5 0 0 1 6.5 11.79l.005.211a4.499 4.499 0 0 1-2.344 3.95l-.156.081V19h16v-2.969l-.156-.08a4.5 4.5 0 0 1-2.34-3.74L17.506 12c0-1.704.947-3.187 2.344-3.95l.156-.082V5Zm-4 4v6h-8V9h8Z"/>`),
		g.Group(children),
	)
}