package mdi

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func RelationOnlyOneToOnlyOne(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<path fill="currentColor" d="M22 15v2h-1v2h-2v-2h-1v2h-2v-2h-5V9H8v2H6V9H5v2H3V9H2V7h1V5h2v2h1V5h2v2h5v8h3v-2h2v2h1v-2h2v2Z"/>`),
		g.Group(children),
	)
}