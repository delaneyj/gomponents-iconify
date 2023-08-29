package material_symbols

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func DeveloperBoardSharp(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<path fill="currentColor" d="M2 21V3h18v4h2v2h-2v2h2v2h-2v2h2v2h-2v4H2Zm4-4h5v-4H6v4Zm6-7h4V7h-4v3Zm-6 2h5V7H6v5Zm6 5h4v-6h-4v6Z"/>`),
		g.Group(children),
	)
}