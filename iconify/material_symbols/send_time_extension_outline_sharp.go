package material_symbols

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func SendTimeExtensionOutlineSharp(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<path fill="currentColor" d="M13 22v-4l4-1l-4-1v-4l10 5l-10 5Zm-4.2-1H3v-5.8q1.2 0 2.1-.762T6 12.5q0-1.175-.9-1.937T3 9.8V4h6q0-1.05.725-1.775T11.5 1.5q1.05 0 1.775.725T14 4h6v9.25l-2-1V6h-6V4q0-.2-.15-.35t-.35-.15q-.2 0-.35.15T11 4v2H5v2.2q1.35.5 2.175 1.675T8 12.5q0 1.425-.825 2.6T5 16.8V19h2.125q.425-1.125 1.45-1.963T11 16.05v2q-1 .2-1.6.938T8.8 21Zm2.7-9.75Z"/>`),
		g.Group(children),
	)
}