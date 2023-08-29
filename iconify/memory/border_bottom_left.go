package memory

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func BorderBottomLeft(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 22 22"),
		g.Raw(`<path fill="currentColor" d="M10 4V2h2v2h-2m8 8v-2h2v2h-2m0 4v-2h2v2h-2M6 4V2h2v2H6m12 4V6h2v2h-2m0-4V2h2v2h-2m-4 0V2h2v2h-2M2 20V2h2v16h16v2H2Z"/>`),
		g.Group(children),
	)
}