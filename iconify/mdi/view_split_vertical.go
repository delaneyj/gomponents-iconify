package mdi

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func ViewSplitVertical(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<path fill="currentColor" d="M13 5h8v14h-8V5M3 5h8v2H3V5m0 6V9h8v2H3m0 8v-2h8v2H3m0-4v-2h8v2H3Z"/>`),
		g.Group(children),
	)
}