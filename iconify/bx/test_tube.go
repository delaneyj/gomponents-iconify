package bx

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func TestTube(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<path fill="currentColor" d="m13.293 2.707l.818.818L3.318 14.318C2.468 15.168 2 16.298 2 17.5s.468 2.332 1.318 3.183C4.169 21.532 5.299 22 6.5 22s2.331-.468 3.182-1.318L20.475 9.889l.818.818l1.414-1.414l-8-8l-1.414 1.414zm3.182 8.354l-2.403-2.404l-1.414 1.414l2.403 2.404l-1.414 1.415l-.99-.99l-1.414 1.414l.99.99l-1.415 1.415l-2.403-2.404L7 15.728l2.403 2.404l-1.136 1.136c-.945.944-2.59.944-3.535 0C4.26 18.795 4 18.168 4 17.5s.26-1.295.732-1.768L15.525 4.939l3.535 3.535l-2.585 2.587z"/>`),
		g.Group(children),
	)
}