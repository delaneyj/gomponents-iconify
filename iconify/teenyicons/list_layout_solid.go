package teenyicons

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func ListLayoutSolid(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 15 15"),
		g.Raw(`<path fill="currentColor" d="M1.5 1A1.5 1.5 0 0 0 0 2.5v2A1.5 1.5 0 0 0 1.5 6h2A1.5 1.5 0 0 0 5 4.5v-2A1.5 1.5 0 0 0 3.5 1h-2ZM7 4h8V3H7v1ZM1.5 9A1.5 1.5 0 0 0 0 10.5v2A1.5 1.5 0 0 0 1.5 14h2A1.5 1.5 0 0 0 5 12.5v-2A1.5 1.5 0 0 0 3.5 9h-2ZM7 12h8v-1H7v1Z"/>`),
		g.Group(children),
	)
}