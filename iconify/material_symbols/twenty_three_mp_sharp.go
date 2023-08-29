package material_symbols

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func TwentyThreeMpSharp(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<path fill="currentColor" d="M6.5 11.5H11V10H8V9h3V5.5H6.5V7h3v1h-3v3.5Zm6.5 0h4.5v-6H13V7h3v1h-2v1h2v1h-3v1.5Zm-7 7h1.5V14h1v3H10v-3h1v4.5h1.5v-6H6v6Zm7.5 0H15V17h3v-4.5h-4.5v6Zm1.5-3V14h1.5v1.5H15ZM3 21V3h18v18H3Z"/>`),
		g.Group(children),
	)
}