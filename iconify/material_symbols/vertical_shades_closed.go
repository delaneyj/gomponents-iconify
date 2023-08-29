package material_symbols

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func VerticalShadesClosed(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<path fill="currentColor" d="M2 21v-2h2V3h16v16h2v2H2Zm4-2h1.5V5H6v14Zm3.5 0H11V5H9.5v14Zm3.5 0h1.5V5H13v14Zm3.5 0H18V5h-1.5v14Z"/>`),
		g.Group(children),
	)
}