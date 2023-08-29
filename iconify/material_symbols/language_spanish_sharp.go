package material_symbols

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func LanguageSpanishSharp(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<path fill="currentColor" d="M13 17v-3h2v1h3v-2h-5V7h7v3h-2V9h-3v2h5v6h-7ZM4 7h7v2H6v2h4v2H6v2h5v2H4V7Z"/>`),
		g.Group(children),
	)
}