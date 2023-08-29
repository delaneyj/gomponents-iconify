package material_symbols

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func Resize(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<path fill="currentColor" d="M3 13v-2h2v2H3Zm0-4V7h2v2H3Zm0-4V3h2v2H3Zm4 0V3h2v2H7Zm4 16v-2h2v2h-2Zm0-16V3h2v2h-2Zm4 16v-2h2v2h-2Zm4 0v-2h2v2h-2Zm0-4v-2h2v2h-2Zm0-4v-2h2v2h-2Zm0-4V5h-4V3h6v6h-2ZM3 21v-6h2v4h4v2H3Z"/>`),
		g.Group(children),
	)
}