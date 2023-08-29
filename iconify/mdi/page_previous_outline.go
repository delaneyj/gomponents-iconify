package mdi

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func PagePreviousOutline(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<path fill="currentColor" d="M2 3h17a2 2 0 0 1 2 2v4h-2V5H2v14h17v-4h2v4a2 2 0 0 1-2 2H2a2 2 0 0 1-2-2V5a2 2 0 0 1 2-2m15 12v-2h7v-2h-7V9l-4 3l4 3M4 13h7v-2H4v2m0-4h7V7H4v2m0 8h4v-2H4v2Z"/>`),
		g.Group(children),
	)
}