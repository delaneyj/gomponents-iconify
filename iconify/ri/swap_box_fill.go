package ri

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func SwapBoxFill(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<path fill="currentColor" d="M3.005 3.003h18a1 1 0 0 1 1 1v16a1 1 0 0 1-1 1h-18a1 1 0 0 1-1-1v-16a1 1 0 0 1 1-1Zm12 4v2h-4v2h4v2l3.5-3l-3.5-3Zm-6 10v-2h4v-2h-4v-2l-3.5 3l3.5 3Z"/>`),
		g.Group(children),
	)
}