package uil

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func Triangle(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<path fill="currentColor" d="m21.87 19.29l-9-15.58a1 1 0 0 0-1.74 0l-9 15.58a1 1 0 0 0 0 1a1 1 0 0 0 .87.5h18a1 1 0 0 0 .87-.5a1 1 0 0 0 0-1Zm-17.14-.5L12 6.21l7.27 12.58Z"/>`),
		g.Group(children),
	)
}