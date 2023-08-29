package material_symbols

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func Hourglass(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<path fill="currentColor" d="M4 22v-2h2v-3q0-1.525.713-2.863T8.7 12q-1.275-.8-1.987-2.138T6 7V4H4V2h16v2h-2v3q0 1.525-.713 2.863T15.3 12q1.275.8 1.988 2.138T18 17v3h2v2H4Z"/>`),
		g.Group(children),
	)
}