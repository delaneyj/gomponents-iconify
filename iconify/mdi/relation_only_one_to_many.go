package mdi

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func RelationOnlyOneToMany(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<path fill="currentColor" d="M22 13v6h-1l-2-2h-8V9H8v2H6V9H5v2H3V9H2V7h1V5h2v2h1V5h2v2h5v8h6l2-2Z"/>`),
		g.Group(children),
	)
}