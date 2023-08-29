package mdi

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func TableSplitCell(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<path fill="currentColor" d="M19 14h2v6H3v-6h2v4h14v-4M3 4v6h2V6h14v4h2V4H3m8 7v2H8v2l-3-3l3-3v2h3m5 0V9l3 3l-3 3v-2h-3v-2h3Z"/>`),
		g.Group(children),
	)
}