package mdi

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func Bio(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<path fill="currentColor" d="M17 12h3a2 2 0 0 1 2 2v3a2 2 0 0 1-2 2h-3a2 2 0 0 1-2-2v-3a2 2 0 0 1 2-2m0 2v3h3v-3h-3M2 7h5a2 2 0 0 1 2 2v2a2 2 0 0 1-2 2a2 2 0 0 1 2 2v2a2 2 0 0 1-2 2H2V7m2 2v3h3V9H4m0 8h3v-3H4v3m7-4h2v6h-2v-6m0-4h2v2h-2V9Z"/>`),
		g.Group(children),
	)
}