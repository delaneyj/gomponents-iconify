package mdi

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func ViewSplitHorizontal(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<path fill="currentColor" d="M3 5h18v2H3V5m0 6V9h18v2H3m0 8v-6h18v6H3Z"/>`),
		g.Group(children),
	)
}