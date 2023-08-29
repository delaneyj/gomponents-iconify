package memory

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func BorderTopBottom(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 22 22"),
		g.Raw(`<path fill="currentColor" d="M2 10h2v2H2v-2m16 0h2v2h-2v-2M2 14h2v2H2v-2m0-8h2v2H2V6m16 0h2v2h-2V6m0 8h2v2h-2v-2M2 18h18v2H2v-2M2 4V2h18v2H2Z"/>`),
		g.Group(children),
	)
}