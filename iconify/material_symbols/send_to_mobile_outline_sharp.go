package material_symbols

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func SendToMobileOutlineSharp(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<path fill="currentColor" d="m17 16l-1.4-1.4l1.55-1.6H12v-2h5.15L15.6 9.4L17 8l4 4l-4 4ZM5 23V1h14v6h-2V6H7v12h10v-1h2v6H5Zm2-3v1h10v-1H7ZM7 4h10V3H7v1Zm0 0V3v1Zm0 16v1v-1Z"/>`),
		g.Group(children),
	)
}