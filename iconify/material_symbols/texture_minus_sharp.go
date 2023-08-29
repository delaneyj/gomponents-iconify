package material_symbols

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func TextureMinusSharp(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<path fill="currentColor" d="M4.4 21H3v-1.4L19.6 3h1.425v1.4L4.4 21Zm4.9 0l2.7-2.7V21H9.3Zm4.7-2v-2h8v2h-8ZM3 14.7v-2.8L11.9 3h2.8L3 14.7Zm12.3.3L21 9.3v2.8L18.1 15h-2.8ZM3 7V3h4L3 7Z"/>`),
		g.Group(children),
	)
}