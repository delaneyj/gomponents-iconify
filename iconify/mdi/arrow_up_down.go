package mdi

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func ArrowUpDown(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<path fill="currentColor" d="M17.45 17.55L12 23l-5.45-5.45l1.41-1.41L11 19.17V4.83L7.96 7.86L6.55 6.45L12 1l5.45 5.45l-1.41 1.41L13 4.83v14.34l3.04-3.03l1.41 1.41Z"/>`),
		g.Group(children),
	)
}