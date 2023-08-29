package mdi

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func NumericNinePlusBox(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<path fill="currentColor" d="M21 5v14a2 2 0 0 1-2 2H5a2 2 0 0 1-2-2V5a2 2 0 0 1 2-2h14a2 2 0 0 1 2 2m-2 6h-2V9h-2v2h-2v2h2v2h2v-2h2v-2m-9-4H8a2 2 0 0 0-2 2v2a2 2 0 0 0 2 2h2v2H6v2h4a2 2 0 0 0 2-2V9a2 2 0 0 0-2-2M8 9h2v2H8V9Z"/>`),
		g.Group(children),
	)
}