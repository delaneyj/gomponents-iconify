package carbon

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func FlightSchedule(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 32 32"),
		g.Raw(`<path fill="currentColor" fill-rule="evenodd" d="m25 23l5 2v-2l-5-2.5V18a1 1 0 0 0-2 0v2.5L18 23v2l5-2v3.5L21 28v1l3-1l3 1v-1l-2-1.5Z"/><path fill="currentColor" fill-rule="evenodd" d="M26 4h-4V2h-2v2h-8V2h-2v2H6a2 2 0 0 0-2 2v20a2 2 0 0 0 2 2h10v-2H6V6h4v2h2V6h8v2h2V6h4v7h2V6a2 2 0 0 0-2-2Z"/>`),
		g.Group(children),
	)
}