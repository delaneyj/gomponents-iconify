package mdi

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func ChevronDoubleUp(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<path fill="currentColor" d="M7.41 18.41L6 17l6-6l6 6l-1.41 1.41L12 13.83l-4.59 4.58m0-6L6 11l6-6l6 6l-1.41 1.41L12 7.83l-4.59 4.58Z"/>`),
		g.Group(children),
	)
}