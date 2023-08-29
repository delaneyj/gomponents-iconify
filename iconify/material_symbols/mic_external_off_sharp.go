package material_symbols

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func MicExternalOffSharp(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<path fill="currentColor" d="M9.5 6.65L5.35 2.5q.35-.275.775-.388T7 2q1.25 0 2.125.862T10 5q0 .45-.138.863T9.5 6.65ZM20 17.15l-2-2V4h-4v7.15l-2-2V2h8v15.15Zm.5 6.15L14 16.8V22H6v-4H5L4 8h1.15L.7 3.5l1.4-1.4l19.8 19.8l-1.4 1.4ZM8 20h4v-5.2l-2.45-2.45L9 18H8v2Z"/>`),
		g.Group(children),
	)
}