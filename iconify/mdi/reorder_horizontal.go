package mdi

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func ReorderHorizontal(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<path fill="currentColor" d="M3 15h18v-2H3v2m0 4h18v-2H3v2m0-8h18V9H3v2m0-6v2h18V5H3Z"/>`),
		g.Group(children),
	)
}