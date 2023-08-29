package mdi

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func RelationZeroOrManyToOneOrMany(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<path fill="currentColor" d="m21 13l-2 2v-2h-2v2h-4V7H9.79a2.5 2.5 0 0 0-4.58 0H5L3 5H2v6h1l2-2h.21a2.5 2.5 0 0 0 4.58 0H11v8h6v2h2v-2l2 2h1v-6M7.5 9a1 1 0 1 1 1-1a1 1 0 0 1-1 1Z"/>`),
		g.Group(children),
	)
}