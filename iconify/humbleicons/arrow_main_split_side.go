package humbleicons

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func ArrowMainSplitSide(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M21 17H4.603M21 17l-3-3m3 3l-3 3M4.603 17H3m1.603 0a6 6 0 0 0 5.145-2.913l2.504-4.174A6 6 0 0 1 17.397 7H21m0 0l-3 3m3-3l-3-3"/>`),
		g.Group(children),
	)
}