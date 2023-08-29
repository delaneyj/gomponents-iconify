package mdi

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func NumericNinePlus(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<path fill="currentColor" d="M19 11h-2V9h-2v2h-2v2h2v2h2v-2h2v-2m-9-4H8c-1.1 0-2 .9-2 2v2a2 2 0 0 0 2 2h2v2H6v2h4c1.11 0 2-.89 2-2V9a2 2 0 0 0-2-2m0 4H8V9h2v2Z"/>`),
		g.Group(children),
	)
}