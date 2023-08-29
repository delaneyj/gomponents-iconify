package material_symbols

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func CropFiveFourOutlineSharp(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<path fill="currentColor" d="M3 20V4h18v16H3Zm2-2h14V6H5v12Zm0 0V6v12Z"/>`),
		g.Group(children),
	)
}