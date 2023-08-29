package material_symbols

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func LineStyle(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<path fill="currentColor" d="M3 20v-2h2v2H3Zm0-4v-2h5v2H3Zm0-4v-2h8v2H3Zm0-4V4h18v4H3Zm4 12v-2h2v2H7Zm2.5-4v-2h5v2h-5Zm1.5 4v-2h2v2h-2Zm2-8v-2h8v2h-8Zm2 8v-2h2v2h-2Zm1-4v-2h5v2h-5Zm3 4v-2h2v2h-2Z"/>`),
		g.Group(children),
	)
}