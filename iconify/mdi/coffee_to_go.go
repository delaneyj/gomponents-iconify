package mdi

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func CoffeeToGo(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<path fill="currentColor" d="M3 19v-2h14l-1.74-1.76l1.41-1.41L20.84 18l-4.17 4.17l-1.41-1.41L17 19H3M17 8V5h-2v3h2m0-5a2 2 0 0 1 2 2v3c0 1.11-.89 2-2 2h-2v1a4 4 0 0 1-4 4H7a4 4 0 0 1-4-4V3h14Z"/>`),
		g.Group(children),
	)
}