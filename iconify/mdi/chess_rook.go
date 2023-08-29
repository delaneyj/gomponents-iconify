package mdi

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func ChessRook(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<path fill="currentColor" d="M5 20h14v2H5v-2M17 2v3h-2V2h-2v3h-2V2H9v3H7V2H5v6h2v10h10V8h2V2h-2Z"/>`),
		g.Group(children),
	)
}