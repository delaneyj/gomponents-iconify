package gg

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func ArrowsVAlt(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<path fill="currentColor" d="M9.172 6.455L7.757 5.041L12 .798l4.243 4.243l-1.415 1.414L13 4.627v14.746l1.828-1.828l1.415 1.414L12 23.202l-4.243-4.243l1.415-1.414L11 19.373V4.627L9.172 6.455Z"/>`),
		g.Group(children),
	)
}