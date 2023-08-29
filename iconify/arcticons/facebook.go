package arcticons

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func Facebook(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 48 48"),
		g.Raw(`<path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" d="M24 42.5V18.57a5.07 5.07 0 0 1 5.08-5.07h0c2.49 0 4.05.74 5.12 2.12m-14.5 7.67h10.15"/><path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" d="M7.48 5.5a2 2 0 0 0-2 2h0v33a2 2 0 0 0 2 2h33.04a2 2 0 0 0 2-2h0v-33a2 2 0 0 0-2-2H7.48Z"/>`),
		g.Group(children),
	)
}