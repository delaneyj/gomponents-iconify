package material_symbols

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func TableLampOutlineSharp(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<path fill="currentColor" d="M13 21v-2h8v2h-8ZM5.525 9H9V5H7.275l-1.75 4ZM16 18V7h-5v4H2.475l3.5-8H11v2h7v13h-2ZM5.525 9H9H5.525Z"/>`),
		g.Group(children),
	)
}