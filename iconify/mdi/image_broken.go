package mdi

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func ImageBroken(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<path fill="currentColor" d="M19 3a2 2 0 0 1 2 2v6h-2v2h-2v2h-2v2h-2v2h-2v2H5a2 2 0 0 1-2-2V5a2 2 0 0 1 2-2h14m2 12v4a2 2 0 0 1-2 2h-4v-2h2v-2h2v-2h2m-2-6.5a.5.5 0 0 0-.5-.5h-13a.5.5 0 0 0-.5.5v7a.5.5 0 0 0 .5.5H11v-1h2v-2h2v-2h2V9h2v-.5Z"/>`),
		g.Group(children),
	)
}