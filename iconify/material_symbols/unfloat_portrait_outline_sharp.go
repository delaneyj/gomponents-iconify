package material_symbols

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func UnfloatPortraitOutlineSharp(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<path fill="currentColor" d="M4 22V2h16v11h-2V4H6v16h7v2Zm10.5-10.925L11.425 8H14V6H8v6h2V9.4l3.075 3.1ZM15 22v-7h5v7Zm-3-10Z"/>`),
		g.Group(children),
	)
}