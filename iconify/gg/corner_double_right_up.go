package gg

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func CornerDoubleRightUp(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<g fill="currentColor"><path d="m16.216 9.25l1.415-1.414l-4.243-4.242l-4.243 4.242L10.56 9.25l2.828-2.828l2.828 2.828Z"/><path d="M10.56 13.493L9.145 12.08l4.243-4.243l4.243 4.243l-1.415 1.414l-1.847-1.847v4.76a4 4 0 0 1-4 4h-4v-2h4a2 2 0 0 0 2-2v-4.723l-1.81 1.81Z"/></g>`),
		g.Group(children),
	)
}