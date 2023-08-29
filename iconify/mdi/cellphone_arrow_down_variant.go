package mdi

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func CellphoneArrowDownVariant(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<path fill="currentColor" d="m18 14l5-5l-1.4-1.4l-2.6 2.6V3h-2v7.2l-2.6-2.6L13 9l5 5m1 2v5c0 1.1-.9 2-2 2H7c-1.1 0-2-.9-2-2V3c0-1.1.9-2 2-2h7v4H7v14h10v-3h2Z"/>`),
		g.Group(children),
	)
}