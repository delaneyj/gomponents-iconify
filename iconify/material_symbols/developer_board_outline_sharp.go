package material_symbols

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func DeveloperBoardOutlineSharp(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<path fill="currentColor" d="M2 21V3h18v4h2v2h-2v2h2v2h-2v2h2v2h-2v4H2Zm2-2h14V5H4v14Zm2-2h5v-4H6v4Zm6-7h4V7h-4v3Zm-6 2h5V7H6v5Zm6 5h4v-6h-4v6Zm-8 2V5v14Z"/>`),
		g.Group(children),
	)
}