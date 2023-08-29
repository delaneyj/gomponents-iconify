package mdi

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func Altimeter(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<path fill="currentColor" d="M7 3v2h10V3H7m2 4v2h6V7H9m-7 .96v8.08L6.03 12L2 7.96m20.03 0L18 12l4.03 4.04V7.96M7 11v2h10v-2H7m2 4v2h6v-2H9m-2 4v2h10v-2H7Z"/>`),
		g.Group(children),
	)
}