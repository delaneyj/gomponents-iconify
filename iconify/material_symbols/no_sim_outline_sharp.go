package material_symbols

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func NoSimOutlineSharp(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<path fill="currentColor" d="m20 17.175l-2-2V4h-7.15l-2 2L7.4 4.6L10 2h10v15.175Zm.5 6.125L6 8.8V20h12v-2.025l2 2V22H4V8l.6-.6L.7 3.5l1.425-1.4L21.9 21.875L20.5 23.3Zm-6.975-12.625Zm-1.875 3.8Z"/>`),
		g.Group(children),
	)
}