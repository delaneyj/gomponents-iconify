package material_symbols

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func LanguageGbEnglishSharp(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<path fill="currentColor" d="M15 15h4v-2h-4v2Zm0-4h4V9h-4v2ZM3 17V7h8v2H5v6h4v-2H7v-2h4v6H3Zm10 0V7h7l1 1v3l-1 1l1 1v3l-1 1h-7Z"/>`),
		g.Group(children),
	)
}