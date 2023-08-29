package mdi

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func RelationOneToZeroOrMany(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<path fill="currentColor" d="m21 13l-2 2h-.21a2.5 2.5 0 0 0-4.58 0H13V7H7V5H5v2H2v2h3v2h2V9h4v8h3.21a2.5 2.5 0 0 0 4.58 0H19l2 2h1v-6m-5.5 4a1 1 0 1 1 1-1a1 1 0 0 1-1 1Z"/>`),
		g.Group(children),
	)
}