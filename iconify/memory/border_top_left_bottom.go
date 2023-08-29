package memory

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func BorderTopLeftBottom(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 22 22"),
		g.Raw(`<path fill="currentColor" d="M18 10h2v2h-2v-2m0-4h2v2h-2V6m0 8h2v2h-2v-2m2 4v2H2V2h18v2H4v14h16Z"/>`),
		g.Group(children),
	)
}