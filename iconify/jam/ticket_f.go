package jam

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func TicketF(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<path fill="currentColor" d="M12 9v11H8.83a3.001 3.001 0 0 0-5.66 0H0V9h4a1 1 0 1 0 0-2H0V0h3.17a3.001 3.001 0 0 0 5.66 0H12v7H8a1 1 0 1 0 0 2h4z"/>`),
		g.Group(children),
	)
}