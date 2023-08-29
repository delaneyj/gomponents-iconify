package mdi

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func RelationManyToOneOrMany(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<path fill="currentColor" d="M22 13v6h-1l-2-2v2h-2v-2h-6V9H5l-2 2H2V5h1l2 2h8v8h4v-2h2v2l2-2Z"/>`),
		g.Group(children),
	)
}