package material_symbols

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func LanguageUsColemakSharp(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<path fill="currentColor" d="M14.5 15H18V9h-3.5v6ZM4 17V7h6.5v2H6v6h4.5v2H4Zm8.5 0V7H20v10h-7.5Z"/>`),
		g.Group(children),
	)
}