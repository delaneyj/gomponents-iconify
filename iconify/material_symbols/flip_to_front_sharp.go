package material_symbols

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func FlipToFrontSharp(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<path fill="currentColor" d="M3 13v-2h2v2H3Zm0 4v-2h2v2H3Zm0 4v-2h2v2H3ZM3 9V7h2v2H3Zm12 12v-2h2v2h-2Zm-8-4V3h14v14H7Zm2-2h10V5H9v10Zm2 6v-2h2v2h-2Zm-4 0v-2h2v2H7Z"/>`),
		g.Group(children),
	)
}