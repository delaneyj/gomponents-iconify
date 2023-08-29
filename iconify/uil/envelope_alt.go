package uil

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func EnvelopeAlt(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<path fill="currentColor" d="M19 4H5a3 3 0 0 0-3 3v10a3 3 0 0 0 3 3h14a3 3 0 0 0 3-3V7a3 3 0 0 0-3-3ZM5 6h14a1 1 0 0 1 1 1l-8 4.88L4 7a1 1 0 0 1 1-1Zm15 11a1 1 0 0 1-1 1H5a1 1 0 0 1-1-1V9.28l7.48 4.57a1 1 0 0 0 1 0L20 9.28Z"/>`),
		g.Group(children),
	)
}