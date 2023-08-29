package solar

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func DoubleAltArrowDownBroken(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="1.5" d="m19 11l-7 6l-1.75-1.5M5 11l2.333 2M5 7l7 6l1.75-1.5M19 7l-2.333 2"/>`),
		g.Group(children),
	)
}