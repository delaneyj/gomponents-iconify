package material_symbols

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func CropSevenFiveOutlineSharp(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<path fill="currentColor" d="M3 19V5h18v14H3Zm2-2h14V7H5v10Zm0 0V7v10Z"/>`),
		g.Group(children),
	)
}