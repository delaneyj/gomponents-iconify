package material_symbols

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func ThirteenMpSharp(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<path fill="currentColor" d="M8.5 11.5H10v-6H7V7h1.5v4.5Zm3.5 0h4.5v-6H12V7h3v1h-2v1h2v1h-3v1.5Zm-6 7h1.5V14h1v3H10v-3h1v4.5h1.5v-6H6v6Zm7.5 0H15V17h3v-4.5h-4.5v6Zm1.5-3V14h1.5v1.5H15ZM3 21V3h18v18H3Z"/>`),
		g.Group(children),
	)
}