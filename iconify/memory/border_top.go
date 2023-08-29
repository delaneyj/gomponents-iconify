package memory

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func BorderTop(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 22 22"),
		g.Raw(`<path fill="currentColor" d="M2 10h2v2H2v-2m16 0h2v2h-2v-2m-8 8h2v2h-2v-2m-4 0h2v2H6v-2m-4-4h2v2H2v-2m0 4h2v2H2v-2M2 6h2v2H2V6m16 0h2v2h-2V6m-4 12h2v2h-2v-2m4 0h2v2h-2v-2m0-4h2v2h-2v-2M2 2h18v2H2V2Z"/>`),
		g.Group(children),
	)
}