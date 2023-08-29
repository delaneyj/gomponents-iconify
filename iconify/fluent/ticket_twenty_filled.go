package fluent

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func TicketTwentyFilled(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 20 20"),
		g.Raw(`<path d="M18 6v1.5a.5.5 0 0 1-.5.5a2 2 0 1 0 0 4a.5.5 0 0 1 .5.5V14a1 1 0 0 1-1 1H3a1 1 0 0 1-1-1v-1.5a.5.5 0 0 1 .5-.5a2 2 0 1 0 0-4a.5.5 0 0 1-.5-.5V6a1 1 0 0 1 1-1h14a1 1 0 0 1 1 1z" fill="currentColor" fill-rule="nonzero"/>`),
		g.Group(children),
	)
}