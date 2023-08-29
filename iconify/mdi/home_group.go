package mdi

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func HomeGroup(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<path fill="currentColor" d="M17 16h-2v6h-3v-5H8v5H5v-6H3l7-6l7 6M6 2l4 4H9v3H7V6H5v3H3V6H2l4-4m12 1l5 5h-1v4h-3V9h-2v3h-1.66L14 10.87V8h-1l5-5Z"/>`),
		g.Group(children),
	)
}