package mdi

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func Warehouse(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<path fill="currentColor" d="M6 19h2v2H6v-2m6-16L2 8v13h2v-8h16v8h2V8L12 3m-4 8H4V9h4v2m6 0h-4V9h4v2m6 0h-4V9h4v2M6 15h2v2H6v-2m4 0h2v2h-2v-2m0 4h2v2h-2v-2m4 0h2v2h-2v-2Z"/>`),
		g.Group(children),
	)
}