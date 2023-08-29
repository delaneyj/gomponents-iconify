package material_symbols

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func OpenInPhoneOutlineSharp(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<path fill="currentColor" d="M5 23v-8h2v3h10V6H7v3H5V1h14v22H5Zm2-3v1h10v-1H7Zm3-4l-1.4-1.4l1.55-1.6H2v-2h8.15L8.6 9.4L10 8l4 4l-4 4ZM7 4h10V3H7v1Zm0 0V3v1Zm0 16v1v-1Z"/>`),
		g.Group(children),
	)
}