package material_symbols

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func CropThreeTwoOutlineSharp(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<path fill="currentColor" d="M3 18V6h18v12H3Zm2-2h14V8H5v8Zm0 0V8v8Z"/>`),
		g.Group(children),
	)
}