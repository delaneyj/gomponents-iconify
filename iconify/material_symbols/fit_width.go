package material_symbols

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func FitWidth(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<path fill="currentColor" d="M3 21V3h2v18H3Zm4-8v-2h2v2H7Zm4 0v-2h2v2h-2Zm4 0v-2h2v2h-2Zm4 8V3h2v18h-2Z"/>`),
		g.Group(children),
	)
}