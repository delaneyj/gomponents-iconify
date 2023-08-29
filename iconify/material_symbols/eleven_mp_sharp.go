package material_symbols

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func ElevenMpSharp(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<path fill="currentColor" d="M9.5 11.5H11v-6H8V7h1.5v4.5Zm5 0H16v-6h-3V7h1.5v4.5Zm-8.5 7h1.5V14h1v3H10v-3h1v4.5h1.5v-6H6v6Zm7.5 0H15V17h3v-4.5h-4.5v6Zm1.5-3V14h1.5v1.5H15ZM3 21V3h18v18H3Z"/>`),
		g.Group(children),
	)
}