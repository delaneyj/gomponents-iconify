package gg

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func CornerDoubleUpLeft(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<g fill="currentColor"><path d="M9.25 7.784L7.836 6.369l-4.242 4.243l4.242 4.243L9.25 13.44l-2.828-2.828L9.25 7.784Z"/><path d="m13.493 13.44l-1.414 1.415l-4.243-4.243L12.08 6.37l1.414 1.415l-1.847 1.846h4.76a4 4 0 0 1 4 4v4h-2v-4a2 2 0 0 0-2-2h-4.723l1.81 1.81Z"/></g>`),
		g.Group(children),
	)
}