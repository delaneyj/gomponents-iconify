package material_symbols

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func OneKOutlineSharp(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<path fill="currentColor" d="M12 15h1.5v-2.25L15.25 15h1.825l-2.325-3l2.325-3H15.25l-1.75 2.25V9H12v6Zm-3.5 0H10V9H7v1.5h1.5V15ZM3 21V3h18v18H3Zm2-2h14V5H5v14Zm0 0V5v14Z"/>`),
		g.Group(children),
	)
}