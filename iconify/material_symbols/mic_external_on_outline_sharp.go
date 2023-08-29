package material_symbols

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func MicExternalOnOutlineSharp(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<path fill="currentColor" d="M4.8 7q-.375-.425-.587-.925T4 5q0-1.25.875-2.125T7 2q1.25 0 2.125.875T10 5q0 .575-.213 1.075T9.2 7H4.8ZM6 22v-4H5L4 8h6L9 18H8v2h4V2h8v20h-2V4h-4v18H6Zm.8-6h.4l.6-6H6.2l.6 6Zm.4-6h-1h1.6h-.6Z"/>`),
		g.Group(children),
	)
}