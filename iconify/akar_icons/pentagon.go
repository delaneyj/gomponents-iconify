package akar_icons

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func Pentagon(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<path fill="none" stroke="currentColor" stroke-width="2" d="M10.75 3a2 2 0 0 1 2.5 0l7.63 6.103a2 2 0 0 1 .63 2.246l-3.031 8.334A2 2 0 0 1 16.599 21H7.401a2 2 0 0 1-1.88-1.317l-3.03-8.334a2 2 0 0 1 .63-2.246L10.75 3Z"/>`),
		g.Group(children),
	)
}