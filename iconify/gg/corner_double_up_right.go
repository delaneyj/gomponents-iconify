package gg

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func CornerDoubleUpRight(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<g fill="currentColor"><path d="m14.75 7.784l1.414-1.415l4.242 4.243l-4.242 4.243l-1.415-1.415l2.829-2.828l-2.829-2.828Z"/><path d="m10.507 13.44l1.414 1.415l4.243-4.243L11.92 6.37l-1.414 1.415l1.847 1.846h-4.76a4 4 0 0 0-4 4v4h2v-4a2 2 0 0 1 2-2h4.723l-1.81 1.81Z"/></g>`),
		g.Group(children),
	)
}