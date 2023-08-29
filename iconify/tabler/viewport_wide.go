package tabler

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func ViewportWide(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M10 12H3l3-3m0 6l-3-3m11 0h7l-3-3m0 6l3-3M3 6V3h18v3M3 18v3h18v-3"/>`),
		g.Group(children),
	)
}