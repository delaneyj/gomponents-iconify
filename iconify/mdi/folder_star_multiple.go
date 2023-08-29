package mdi

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func FolderStarMultiple(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<path fill="currentColor" d="M2 6H0v14c0 1.11.895 2 2 2h18v-2H2V6m20-2h-8l-2-2H6c-1.11 0-2 .89-2 2v12a2 2 0 0 0 2 2h16a2 2 0 0 0 2-2V6c0-1.11-.89-2-2-2m-2.06 11L17 13.27L14.06 15l.78-3.34l-2.59-2.24l3.41-.29L17 6l1.34 3.13l3.41.29l-2.59 2.24l.78 3.34Z"/>`),
		g.Group(children),
	)
}