package mdi

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func RelationOneOrManyToOnlyOne(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<path fill="currentColor" d="M22 15v2h-1v2h-2v-2h-1v2h-2v-2h-5V9H7v2H5V9l-2 2H2V5h1l2 2V5h2v2h6v8h3v-2h2v2h1v-2h2v2Z"/>`),
		g.Group(children),
	)
}