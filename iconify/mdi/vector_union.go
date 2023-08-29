package mdi

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func VectorUnion(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<path fill="currentColor" d="M3 1c-1.11 0-2 .89-2 2v11c0 1.11.89 2 2 2h4v4c0 1.11.89 2 2 2h11c1.11 0 2-.89 2-2V9c0-1.11-.89-2-2-2h-4V3c0-1.11-.89-2-2-2H3m0 2h11v6h6v11H9v-6H3V3Z"/>`),
		g.Group(children),
	)
}