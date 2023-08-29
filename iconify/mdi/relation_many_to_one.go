package mdi

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func RelationManyToOne(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<path fill="currentColor" d="M22 15v2h-3v2h-2v-2h-6V9H5l-2 2H2V5h1l2 2h8v8h4v-2h2v2Z"/>`),
		g.Group(children),
	)
}