package mdi

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func CoffeeOutline(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<path fill="currentColor" d="M2 21v-2h18v2H2M20 8V5h-2v3h2m0-5a2 2 0 0 1 2 2v3a2 2 0 0 1-2 2h-2v3a4 4 0 0 1-4 4H8a4 4 0 0 1-4-4V3h16m-4 2H6v8a2 2 0 0 0 2 2h6a2 2 0 0 0 2-2V5Z"/>`),
		g.Group(children),
	)
}