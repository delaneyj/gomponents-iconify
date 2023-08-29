package material_symbols

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func SecurityUpdateGoodOutlineSharp(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<path fill="currentColor" d="m11.05 15l-2.8-2.8l1.4-1.4l1.4 1.4l3.55-3.55l1.4 1.4L11.05 15ZM5 23V1h14v22H5Zm2-3v1h10v-1H7Zm0-2h10V6H7v12ZM7 4h10V3H7v1Zm0 0V3v1Zm0 16v1v-1Z"/>`),
		g.Group(children),
	)
}