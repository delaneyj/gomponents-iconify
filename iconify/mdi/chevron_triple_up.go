package mdi

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func ChevronTripleUp(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<path fill="currentColor" d="M16.59 9.42L12 4.83L7.41 9.42L6 8l6-6l6 6l-1.41 1.42m0 6L12 10.83l-4.59 4.59L6 14l6-6l6 6l-1.41 1.42m0 6L12 16.83l-4.59 4.59L6 20l6-6l6 6l-1.41 1.42Z"/>`),
		g.Group(children),
	)
}