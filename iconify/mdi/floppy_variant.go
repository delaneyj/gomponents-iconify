package mdi

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func FloppyVariant(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<path fill="currentColor" d="M3 3v18h18V3H3m9 7a2 2 0 0 1 2 2a2 2 0 0 1-2 2a2 2 0 0 1-2-2a2 2 0 0 1 2-2m0 5a1 1 0 0 1 1 1v3a1 1 0 0 1-1 1a1 1 0 0 1-1-1v-3a1 1 0 0 1 1-1Z"/>`),
		g.Group(children),
	)
}