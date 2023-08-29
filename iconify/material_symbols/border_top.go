package material_symbols

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func BorderTop(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<path fill="currentColor" d="M3 21v-2h2v2H3Zm0-4v-2h2v2H3Zm0-4v-2h2v2H3Zm0-4V7h2v2H3Zm0-4V3h18v2H3Zm4 16v-2h2v2H7Zm0-8v-2h2v2H7Zm4 8v-2h2v2h-2Zm0-4v-2h2v2h-2Zm0-4v-2h2v2h-2Zm0-4V7h2v2h-2Zm4 12v-2h2v2h-2Zm0-8v-2h2v2h-2Zm4 8v-2h2v2h-2Zm0-4v-2h2v2h-2Zm0-4v-2h2v2h-2Zm0-4V7h2v2h-2Z"/>`),
		g.Group(children),
	)
}