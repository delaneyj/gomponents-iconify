package uim

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func CornerUpLeft(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<path fill="currentColor" d="M14.686 6.628H7.729l2.92-2.921a1 1 0 0 0-1.413-1.414L4.608 6.921a1.003 1.003 0 0 0 0 1.415l4.628 4.628a1 1 0 0 0 1.414-1.414L7.728 8.628h6.958a3.003 3.003 0 0 1 3 3V21a1 1 0 0 0 2 0v-9.372a5.006 5.006 0 0 0-5-5Z"/>`),
		g.Group(children),
	)
}