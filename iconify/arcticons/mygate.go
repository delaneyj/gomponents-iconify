package arcticons

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func Mygate(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 48 48"),
		g.Raw(`<path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" d="M42.5 40.5v-33a2 2 0 0 0-2-2h-33a2 2 0 0 0-2 2v33a2 2 0 0 0 2 2h11.876v-9.25A4.624 4.624 0 0 1 24 37.874h0a4.624 4.624 0 0 1 4.624-4.624v9.25H40.5a2 2 0 0 0 2-2v-.418"/>`),
		g.Group(children),
	)
}