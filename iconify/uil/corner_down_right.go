package uil

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func CornerDownRight(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<path fill="currentColor" d="M19.61 16a1.19 1.19 0 0 0-.22-.33L14.76 11a1 1 0 0 0-1.41 1.41l2.92 2.92h-7a3 3 0 0 1-3-3V3a1 1 0 1 0-2 0v9.37a5 5 0 0 0 5 5h7l-2.92 2.92a1 1 0 0 0 0 1.42a1 1 0 0 0 .71.29a1 1 0 0 0 .7-.29l4.63-4.63a1 1 0 0 0 .22-.33a1 1 0 0 0 0-.75Z"/>`),
		g.Group(children),
	)
}