package tabler

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func ViewportNarrow(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M3 12h7L7 9m0 6l3-3m11 0h-7l3-3m0 6l-3-3M9 6V3h6v3M9 18v3h6v-3"/>`),
		g.Group(children),
	)
}