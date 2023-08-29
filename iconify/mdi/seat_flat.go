package mdi

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func SeatFlat(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<path fill="currentColor" d="M22 11v2H9V7h9a4 4 0 0 1 4 4M2 14v2h6v2h8v-2h6v-2M7.14 12.1C8.3 10.91 8.28 9 7.1 7.86C5.91 6.7 4 6.72 2.86 7.9C1.7 9.09 1.72 11 2.9 12.14c1.19 1.16 3.1 1.14 4.24-.04Z"/>`),
		g.Group(children),
	)
}