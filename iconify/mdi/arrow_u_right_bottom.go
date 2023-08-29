package mdi

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func ArrowURightBottom(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<path fill="currentColor" d="m20 16l-5.5 5.5l-1.42-1.41L16.17 17H10.5a6.5 6.5 0 1 1 0-13H18v2h-7.5C8 6 6 8 6 10.5S8 15 10.5 15h5.67l-3.08-3.09l1.41-1.41L20 16Z"/>`),
		g.Group(children),
	)
}