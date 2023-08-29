package material_symbols

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func Crossword(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<path fill="currentColor" d="M9 22v-6h6v6H9Zm-7-7V9h6v6H2Zm7 0V9h6v6H9Zm7 0V9h6v6h-6Zm0-7V2h6v6h-6Z"/>`),
		g.Group(children),
	)
}