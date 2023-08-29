package mdi

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func ChevronTripleLeft(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<path fill="currentColor" d="M9.42 7.41L4.83 12l4.59 4.59L8 18l-6-6l6-6l1.42 1.41m6 0L10.83 12l4.59 4.59L14 18l-6-6l6-6l1.42 1.41m6 0L16.83 12l4.59 4.59L20 18l-6-6l6-6l1.42 1.41Z"/>`),
		g.Group(children),
	)
}