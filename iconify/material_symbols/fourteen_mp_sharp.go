package material_symbols

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func FourteenMpSharp(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<path fill="currentColor" d="M8.5 11.5H10v-6H7V7h1.5v4.5Zm6 0H16V10h1V8.5h-1v-3h-1.5v3H13v-3h-1.5V10h3v1.5Zm-8.5 7h1.5V14h1v3H10v-3h1v4.5h1.5v-6H6v6Zm7.5 0H15V17h3v-4.5h-4.5v6Zm1.5-3V14h1.5v1.5H15ZM3 21V3h18v18H3Z"/>`),
		g.Group(children),
	)
}