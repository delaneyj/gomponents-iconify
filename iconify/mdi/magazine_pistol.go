package mdi

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func MagazinePistol(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<path fill="currentColor" d="m14 1l-2 2H7l2 18H8v2h10v-2L16 1M9 5h3l.24 2h-3m.23 2h3l.24 2h-3m.23 2h3l.24 2h-3m.23 2h3l.24 2h-3Z"/>`),
		g.Group(children),
	)
}