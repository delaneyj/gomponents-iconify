package memory

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func BorderOutside(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 22 22"),
		g.Raw(`<path fill="currentColor" d="M2 2h18v18H2V2m2 2v14h14V4H4m6 2h2v2h-2V6m0 4h2v2h-2v-2m-4 0h2v2H6v-2m8 0h2v2h-2v-2m-4 4h2v2h-2v-2Z"/>`),
		g.Group(children),
	)
}