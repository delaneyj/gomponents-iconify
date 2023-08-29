package material_symbols

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func BorderRight(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<path fill="currentColor" d="M3 5V3h2v2H3Zm4 0V3h2v2H7Zm4 0V3h2v2h-2Zm4 0V3h2v2h-2Zm4 16V3h2v18h-2ZM3 9V7h2v2H3Zm8 0V7h2v2h-2Zm-8 4v-2h2v2H3Zm4 0v-2h2v2H7Zm4 0v-2h2v2h-2Zm4 0v-2h2v2h-2ZM3 17v-2h2v2H3Zm8 0v-2h2v2h-2Zm-8 4v-2h2v2H3Zm4 0v-2h2v2H7Zm4 0v-2h2v2h-2Zm4 0v-2h2v2h-2Z"/>`),
		g.Group(children),
	)
}