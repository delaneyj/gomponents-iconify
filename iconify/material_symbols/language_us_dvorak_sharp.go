package material_symbols

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func LanguageUsDvorakSharp(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<path fill="currentColor" d="M6 15h3.25V9H6v6Zm-2 2V7h6.25l1 1v8l-1 1H4Zm11.625 0L12.25 7h2l2.375 6.95L19 7h2l-3.375 10h-2Z"/>`),
		g.Group(children),
	)
}