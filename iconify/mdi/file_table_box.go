package mdi

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func FileTableBox(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<path fill="currentColor" d="M19 3H5c-1.11 0-2 .89-2 2v14c0 1.1.9 2 2 2h14c1.1 0 2-.9 2-2V5a2 2 0 0 0-2-2M9 18H6v-2h3v2m0-3H6v-2h3v2m0-3H6v-2h3v2m4 6h-3v-2h3v2m0-3h-3v-2h3v2m0-3h-3v-2h3v2Z"/>`),
		g.Group(children),
	)
}