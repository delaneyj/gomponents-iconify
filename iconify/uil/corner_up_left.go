package uil

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func CornerUpLeft(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<path fill="currentColor" d="M14.69 6.63h-7l2.92-2.92a1 1 0 0 0 0-1.42a1 1 0 0 0-1.41 0L4.61 6.92a1 1 0 0 0-.22.33a1 1 0 0 0 0 .76a1.19 1.19 0 0 0 .22.33L9.24 13a1 1 0 0 0 .7.3a1 1 0 0 0 .71-1.71L7.73 8.63h7a3 3 0 0 1 3 3V21a1 1 0 0 0 2 0v-9.37a5 5 0 0 0-5.04-5Z"/>`),
		g.Group(children),
	)
}