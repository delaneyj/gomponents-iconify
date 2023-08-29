package gg

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func CornerDoubleLeftUp(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<g fill="currentColor"><path d="M7.784 9.25L6.37 7.836l4.242-4.242l4.243 4.242L13.44 9.25l-2.829-2.828L7.784 9.25Z"/><path d="m13.44 13.493l1.415-1.414l-4.243-4.243L6.37 12.08l1.414 1.414l1.847-1.847v4.76a4 4 0 0 0 4 4h4v-2h-4a2 2 0 0 1-2-2v-4.723l1.81 1.81Z"/></g>`),
		g.Group(children),
	)
}