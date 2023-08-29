package material_symbols

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func InventoryTwoSharp(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<path fill="currentColor" d="M3 22V9H2V2h20v7h-1v13H3ZM4 7h16V4H4v3Zm5 7h6v-2H9v2Z"/>`),
		g.Group(children),
	)
}