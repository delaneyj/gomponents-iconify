package material_symbols

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func TextSelectEnd(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<path fill="currentColor" d="M11 5V3h2v2h-2Zm0 16v-2h2v2h-2ZM7 5V3h2v2H7Zm0 16v-2h2v2H7ZM3 5V3h2v2H3Zm0 4V7h2v2H3Zm0 4v-2h2v2H3Zm0 4v-2h2v2H3Zm0 4v-2h2v2H3Zm12 0v-2h2V5h-2V3h6v2h-2v14h2v2h-6Z"/>`),
		g.Group(children),
	)
}