package material_symbols

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func HighDensity(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<path fill="currentColor" d="M3 20.975v-2h2v2H3ZM7 21v-2h2v2H7Zm4 0v-2h2v2h-2Zm4 0v-2h2v2h-2Zm4 0v-2h2v2h-2Zm0-4v-2h2v2h-2Zm0-4v-2h2v2h-2Zm0-4V7h2v2h-2Zm0-3.975v-2h2v2h-2ZM3 17V3h14v14H3Z"/>`),
		g.Group(children),
	)
}