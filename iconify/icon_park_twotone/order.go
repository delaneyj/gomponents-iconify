package icon_park_twotone

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func Order(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 48 48"),
		g.Raw(`<mask id="ipTOrder0"><g fill="none" stroke="#fff" stroke-linejoin="round" stroke-width="4"><path fill="#555" d="M33.05 7H38a2 2 0 0 1 2 2v33a2 2 0 0 1-2 2H10a2 2 0 0 1-2-2V9a2 2 0 0 1 2-2h7v3h14V7h2.05Z"/><path stroke-linecap="round" d="M17 4h14v6H17zm10 15l-8 8.001h10.004l-8.004 8"/></g></mask><path fill="currentColor" d="M0 0h48v48H0z" mask="url(#ipTOrder0)"/>`),
		g.Group(children),
	)
}