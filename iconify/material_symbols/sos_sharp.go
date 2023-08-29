package material_symbols

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func SosSharp(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<path fill="currentColor" d="M8.5 17V7h7v10h-7ZM1 17v-2h4v-2H1V7h6v2H3v2h4v6H1Zm16 0v-2h4v-2h-4V7h6v2h-4v2h4v6h-6Zm-6.5-2h3V9h-3v6Z"/>`),
		g.Group(children),
	)
}