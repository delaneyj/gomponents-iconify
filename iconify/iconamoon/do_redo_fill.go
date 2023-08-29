package iconamoon

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func DoRedoFill(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<path fill="currentColor" fill-rule="evenodd" d="M15.617 5.076a1 1 0 0 1 1.09.217l3 3a1 1 0 0 1 0 1.414l-3 3A1 1 0 0 1 15 12v-2H8a3 3 0 1 0 0 6h11a1 1 0 1 1 0 2H8A5 5 0 0 1 8 8h7V6a1 1 0 0 1 .617-.924Z" clip-rule="evenodd"/>`),
		g.Group(children),
	)
}