package mdi

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func DragHorizontal(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<path fill="currentColor" d="M3 15v-2h2v2H3m0-4V9h2v2H3m4 4v-2h2v2H7m0-4V9h2v2H7m4 4v-2h2v2h-2m0-4V9h2v2h-2m4 4v-2h2v2h-2m0-4V9h2v2h-2m4 4v-2h2v2h-2m0-4V9h2v2h-2Z"/>`),
		g.Group(children),
	)
}