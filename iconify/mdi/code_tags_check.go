package mdi

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func CodeTagsCheck(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<path fill="currentColor" d="M6.59 3.41L2 8l4.59 4.6L8 11.18L4.82 8L8 4.82L6.59 3.41m5.82 0L11 4.82L14.18 8L11 11.18l1.41 1.42L17 8l-4.59-4.59m9.18 8.18l-8.09 8.09L9.83 16l-1.41 1.41l5.08 5.09L23 13l-1.41-1.41Z"/>`),
		g.Group(children),
	)
}