package mdi

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func RelationOneOrManyToMany(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<path fill="currentColor" d="M22 13v6h-1l-2-2h-8V9H7v2H5V9l-2 2H2V5h1l2 2V5h2v2h6v8h6l2-2Z"/>`),
		g.Group(children),
	)
}