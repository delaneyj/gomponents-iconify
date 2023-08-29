package material_symbols

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func BreakingNewsOutlineSharp(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<path fill="currentColor" d="M6 17h2v-2H6v2Zm0-4h2V7H6v6Zm5 4h7v-2h-7v2Zm0-4h7v-2h-7v2Zm0-4h7V7h-7v2ZM2 21V3h20v18H2Zm2-2h16V5H4v14Zm0 0V5v14Z"/>`),
		g.Group(children),
	)
}