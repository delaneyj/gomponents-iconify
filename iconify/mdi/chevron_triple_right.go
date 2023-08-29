package mdi

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func ChevronTripleRight(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<path fill="currentColor" d="M14.58 16.59L19.17 12l-4.59-4.59L16 6l6 6l-6 6l-1.42-1.41m-6 0L13.17 12L8.58 7.41L10 6l6 6l-6 6l-1.42-1.41m-6 0L7.17 12L2.58 7.41L4 6l6 6l-6 6l-1.42-1.41Z"/>`),
		g.Group(children),
	)
}