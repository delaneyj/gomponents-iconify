package mdi

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func IronBoard(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<path fill="currentColor" d="M19 6h-3V4.8c0-1.12.56-2.17 1.5-2.8c.94.63 1.5 1.68 1.5 2.8V6m-3.72 7.63L19.59 19c.79.06 1.41.7 1.41 1.5c0 .83-.67 1.5-1.5 1.5s-1.5-.67-1.5-1.5v-.24l-4-5.03l-4 5.03v.24c0 .83-.67 1.5-1.5 1.5S7 21.33 7 20.5c0-.8.63-1.44 1.41-1.5l4.31-5.37L9 9H3V7h18v2h-2l-3.72 4.63M14 12l2.42-3h-4.84L14 12Z"/>`),
		g.Group(children),
	)
}