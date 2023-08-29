package material_symbols

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func FourKPlusSharp(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<path fill="currentColor" d="M16.5 14h1v-1.5H19v-1h-1.5V10h-1v1.5H15v1h1.5V14Zm-5 1H13v-2.25L14.75 15h1.75l-2.25-3l2.25-3h-1.75L13 11.25V9h-1.5v6Zm-3 0H10v-1.5h1V12h-1V9H8.5v3H7V9H5.5v4.5h3V15ZM3 21V3h18v18H3Z"/>`),
		g.Group(children),
	)
}