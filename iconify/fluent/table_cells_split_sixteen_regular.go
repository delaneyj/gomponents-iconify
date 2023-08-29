package fluent

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func TableCellsSplitSixteenRegular(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 20 20"),
		g.Raw(`<path fill="currentColor" d="M2 4.5A2.5 2.5 0 0 1 4.5 2h7A2.5 2.5 0 0 1 14 4.5v7a2.5 2.5 0 0 1-2.5 2.5h-7A2.5 2.5 0 0 1 2 11.5v-7ZM3.085 4H7V3H4.5a1.5 1.5 0 0 0-1.415 1ZM8 3v1h4.915A1.5 1.5 0 0 0 11.5 3H8Zm5 2H3v6h10V5Zm-.085 7H8v1h3.5a1.5 1.5 0 0 0 1.415-1ZM7 13v-1H3.085A1.5 1.5 0 0 0 4.5 13H7Zm0-3V6h1v4H7Z"/>`),
		g.Group(children),
	)
}