package material_symbols

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func BakeryDiningSharp(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<path fill="currentColor" d="m20.45 18.2l-1.95-.95l1.825-4.5l2.2 4.25l-2.075 1.2ZM15.5 17l.85-9.425l3.8 1.525L17 17h-1.5ZM7 17L3.85 9.1l3.8-1.525L8.5 17H7Zm-3.45 1.2L1.475 17l2.2-4.25l1.825 4.5l-1.95.95ZM10 17L9 6h6l-1 11h-4Z"/>`),
		g.Group(children),
	)
}