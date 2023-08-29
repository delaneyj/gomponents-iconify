package material_symbols

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func TableLampSharp(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<path fill="currentColor" d="M13 21v-2h8v2h-8Zm3-3V7h-5v4H2.475l3.5-8H11v2h7v13h-2Z"/>`),
		g.Group(children),
	)
}