package majesticons

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func TShirtLine(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M3 10V5a1 1 0 0 1 1-1h4c0 1 .8 3 4 3s4-2 4-3h4a1 1 0 0 1 1 1v5a1 1 0 0 1-1 1h-1a1 1 0 0 0-1 1v7a1 1 0 0 1-1 1H7a1 1 0 0 1-1-1v-7a1 1 0 0 0-1-1H4a1 1 0 0 1-1-1z"/>`),
		g.Group(children),
	)
}