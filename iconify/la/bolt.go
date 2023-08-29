package la

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func Bolt(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 32 32"),
		g.Raw(`<path fill="currentColor" d="M16 6v9l4 1l-3.898 10H16v-8.031l-4-1L15.898 6H16m2-2h-3.512l-.472 1.328l-3.903 10.973l-.734 2.074l2.137.535l2.484.621V28h3.469l.496-1.273l3.898-10l.825-2.118L18 13.438z"/>`),
		g.Group(children),
	)
}