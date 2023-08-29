package arcticons

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func Oitrack(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 48 48"),
		g.Raw(`<path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" d="M40.5 5.5h-33a2 2 0 0 0-2 2v33a2 2 0 0 0 2 2h33a2 2 0 0 0 2-2v-33a2 2 0 0 0-2-2Z"/><path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" d="m15.53 18.114l15.952-7.216a1 1 0 0 1 1.323.5l2.266 5.01a1 1 0 0 1-.499 1.322l-2.909 1.316a1 1 0 0 0-.499 1.324l4.52 9.993a1 1 0 0 1-.499 1.323l-5.578 2.524a1 1 0 0 1-1.323-.499l-4.52-9.993a1 1 0 0 0-1.324-.5l-3.82 1.729M17.039 35.13a4.58 4.58 0 0 0-4.128-4.15l-.012.012m.736 3.387a1.836 1.836 0 1 0 0 2.597h0a1.837 1.837 0 0 0 0-2.597Zm7.624.48a9.006 9.006 0 0 0-8.105-8.104"/>`),
		g.Group(children),
	)
}