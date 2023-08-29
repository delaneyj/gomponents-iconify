package mdi

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func ArrowDownRight(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<path fill="currentColor" d="M21.5 14.5L16 20l-5.5-5.5l1.41-1.41L15 16.17V10.5C15 8 13 6 10.5 6H4V4h6.5a6.5 6.5 0 0 1 6.5 6.5v5.67l3.09-3.09l1.41 1.42Z"/>`),
		g.Group(children),
	)
}