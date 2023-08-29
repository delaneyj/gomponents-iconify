package material_symbols

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func FitPage(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<path fill="currentColor" d="M3 21V3h2v18H3Zm8 0v-2h2v2h-2Zm0-4v-2h2v2h-2Zm-4-4v-2h2v2H7Zm4 0v-2h2v2h-2Zm4 0v-2h2v2h-2Zm-4-4V7h2v2h-2Zm0-4V3h2v2h-2Zm8 16V3h2v18h-2Z"/>`),
		g.Group(children),
	)
}