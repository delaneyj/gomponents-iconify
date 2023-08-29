package mdi

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func HomeModern(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<path fill="currentColor" d="M6 21V8a2 2 0 0 1 2-2l8-3v3a2 2 0 0 1 2 2v13h-6v-5H8v5H6m8-2h2v-3h-2v3m-6-6h2V9H8v4m4 0h4V9h-4v4Z"/>`),
		g.Group(children),
	)
}