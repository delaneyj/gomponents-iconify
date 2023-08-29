package material_symbols

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func CropNineSixteenOutlineSharp(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<path fill="currentColor" d="M7 21V3h10v18H7ZM9 5v14h6V5H9Zm0 14V5v14Z"/>`),
		g.Group(children),
	)
}