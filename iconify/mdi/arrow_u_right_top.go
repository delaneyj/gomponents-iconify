package mdi

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func ArrowURightTop(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<path fill="currentColor" d="M10.5 18H18v2h-7.5a6.5 6.5 0 1 1 0-13h5.67l-3.09-3.09L14.5 2.5L20 8l-5.5 5.5l-1.41-1.41L16.17 9H10.5C8 9 6 11 6 13.5S8 18 10.5 18Z"/>`),
		g.Group(children),
	)
}