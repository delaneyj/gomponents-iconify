package material_symbols

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func FlipSharp(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<path fill="currentColor" d="M9 21H3V3h6v2H5v14h4v2Zm2 2V1h2v22h-2Zm4-2v-2h2v2h-2Zm0-16V3h2v2h-2Zm4 12v-2h2v2h-2Zm0-4v-2h2v2h-2Zm0-4V7h2v2h-2Zm2-4h-2V3h2v2Zm-2 16v-2h2v2h-2Z"/>`),
		g.Group(children),
	)
}