package iconamoon

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func ScreenNormalThin(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" d="M9 4v3a2 2 0 0 1-2 2H4m11 11v-3a2 2 0 0 1 2-2h3m0-6h-3a2 2 0 0 1-2-2V4M4 15h3a2 2 0 0 1 2 2v3"/>`),
		g.Group(children),
	)
}