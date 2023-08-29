package material_symbols

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func UnfloatPortraitSharp(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<path fill="currentColor" d="m13.075 12.5l1.425-1.425L11.425 8H14V6H8v6h2V9.4ZM15 22v-7h5v7ZM4 22V2h16v11h-7v9Z"/>`),
		g.Group(children),
	)
}