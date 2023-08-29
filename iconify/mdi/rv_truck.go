package mdi

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func RvTruck(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<path fill="currentColor" d="M20 8h2V6l-2-2H3a2 2 0 0 0-2 2v9a2 2 0 0 0 2 2h1a3 3 0 0 0 3 3a3 3 0 0 0 3-3h5a3 3 0 0 0 6 0h2v-5M7 18.5a1.5 1.5 0 0 1 0-3a1.5 1.5 0 0 1 0 3M9 12H3V9h6m5 6h-3V9h3m4 9.5a1.5 1.5 0 1 1 1.5-1.5a1.54 1.54 0 0 1-1.5 1.5M17 12V9.5h2.5l2 2.5Z"/>`),
		g.Group(children),
	)
}