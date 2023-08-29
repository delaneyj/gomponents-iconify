package memory

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func BorderTopRightBottom(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 22 22"),
		g.Raw(`<path fill="currentColor" d="M4 12H2v-2h2v2m0 4H2v-2h2v2m0-8H2V6h2v2M2 4V2h18v18H2v-2h16V4H2Z"/>`),
		g.Group(children),
	)
}