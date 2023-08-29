package mdi

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func FridgeTop(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<path fill="currentColor" d="M8 8V6h2v2H8M7 2h10a2 2 0 0 1 2 2v15a2 2 0 0 1-2 2v1h-2v-1H9v1H7v-1a2 2 0 0 1-2-2V4a2 2 0 0 1 2-2m0 2v5h10V4H7m1 8v3h2v-3H8Z"/>`),
		g.Group(children),
	)
}