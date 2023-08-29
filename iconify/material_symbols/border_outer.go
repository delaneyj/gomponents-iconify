package material_symbols

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func BorderOuter(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<path fill="currentColor" d="M7 13v-2h2v2H7Zm4 4v-2h2v2h-2Zm0-4v-2h2v2h-2Zm0-4V7h2v2h-2Zm4 4v-2h2v2h-2ZM5 19h14V5H5v14Zm-2 2V3h18v18H3Z"/>`),
		g.Group(children),
	)
}