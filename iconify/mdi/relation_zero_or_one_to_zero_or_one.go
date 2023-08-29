package mdi

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func RelationZeroOrOneToZeroOrOne(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<path fill="currentColor" d="M21 15v-2h-2v2h-.21a2.5 2.5 0 0 0-4.58 0H13V7H9.79a2.5 2.5 0 0 0-4.58 0H5V5H3v2H2v2h1v2h2V9h.21a2.5 2.5 0 0 0 4.58 0H11v8h3.21a2.5 2.5 0 0 0 4.58 0H19v2h2v-2h1v-2M7.5 9a1 1 0 1 1 1-1a1 1 0 0 1-1 1m9 8a1 1 0 1 1 1-1a1 1 0 0 1-1 1Z"/>`),
		g.Group(children),
	)
}