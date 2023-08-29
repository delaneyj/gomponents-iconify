package mdi

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func CardMultipleOutline(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<path fill="currentColor" d="M21 16V6H7v10h14m0-12c.53 0 1.04.21 1.41.59c.38.37.59.88.59 1.41v10c0 .53-.21 1.04-.59 1.41c-.37.38-.88.59-1.41.59H7a2 2 0 0 1-2-2V6c0-1.11.89-2 2-2h14M3 20h15v2H3c-.53 0-1.04-.21-1.41-.59C1.21 21.04 1 20.53 1 20V9h2v11Z"/>`),
		g.Group(children),
	)
}