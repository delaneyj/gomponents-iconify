package mdi

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func Identifier(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<path fill="currentColor" d="M10 7v2H9v6h1v2H6v-2h1V9H6V7h4m6 0a2 2 0 0 1 2 2v6c0 1.11-.89 2-2 2h-4V7m4 2h-2v6h2V9Z"/>`),
		g.Group(children),
	)
}