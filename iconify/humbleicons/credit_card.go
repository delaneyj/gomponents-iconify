package humbleicons

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func CreditCard(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M4 10V8a2 2 0 0 1 2-2h12a2 2 0 0 1 2 2v2M4 10h16M4 10v6a2 2 0 0 0 2 2h12a2 2 0 0 0 2-2v-6M7 15h5"/>`),
		g.Group(children),
	)
}