package material_symbols

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func FullCoverageOutlineSharp(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<path fill="currentColor" d="M2 21V7h2v12h15v2H2Zm4-4V3h17v14H6Zm2-2h13V5H8v10Zm2-3h4V7h-4v5Zm5 0h4v-2h-4v2Zm0-3h4V7h-4v2Zm-7 6V5v10Z"/>`),
		g.Group(children),
	)
}