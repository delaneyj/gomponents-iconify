package material_symbols

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func PhonelinkEraseOutlineSharp(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<path fill="currentColor" d="M4 23V1h14v6h-2V6H6v12h10v-1h2v6H4Zm2-3v1h10v-1H6ZM6 4h10V3H6v1Zm0 0V3v1Zm0 16v1v-1Zm8.4-4L13 14.6l2.6-2.6L13 9.4L14.4 8l2.6 2.6L19.6 8L21 9.4L18.4 12l2.6 2.6l-1.4 1.4l-2.6-2.6l-2.6 2.6Z"/>`),
		g.Group(children),
	)
}