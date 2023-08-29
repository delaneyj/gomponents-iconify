package material_symbols

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func LabProfileSharp(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<path fill="currentColor" d="M8 12h8v-2H8v2Zm0-4h8V6H8v2Zm12 12.55L14.975 14H4V2h16v18.55ZM4 22v-6h10l4.6 6H4Z"/>`),
		g.Group(children),
	)
}