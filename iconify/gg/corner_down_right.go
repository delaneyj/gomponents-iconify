package gg

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func CornerDownRight(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<path fill="currentColor" d="M6.85 13.4a2 2 0 0 1-2-2v-8h-2v8a4 4 0 0 0 4 4h10.306l-3.785 3.785l1.415 1.414l6.364-6.364l-6.364-6.364l-1.415 1.415l4.115 4.114H6.85Z"/>`),
		g.Group(children),
	)
}