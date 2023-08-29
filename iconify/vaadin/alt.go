package vaadin

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func Alt(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 0 0"),
		g.Raw(`<path fill="currentColor" d="M3.89 9h2.22L5 5.1L3.89 9z"/><path fill="currentColor" d="M0 0v16h16V0H0zm7 12l-.61-2H3.61L3 12H2l2.27-8h1.46L8 12H7zm3 0H9V3h1v9zm4-5h-1v3.5s0 .5 1 .5v1c-1 0-2-.44-2-1.44V7h-.5V6h.5V5h1v1h1v1z"/>`),
		g.Group(children),
	)
}