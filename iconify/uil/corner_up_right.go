package uil

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func CornerUpRight(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<path fill="currentColor" d="M19.61 7.25a1 1 0 0 0-.22-.33l-4.63-4.63a1 1 0 0 0-1.41 0a1 1 0 0 0 0 1.42l2.92 2.92h-8a4 4 0 0 0-4 4V21a1 1 0 0 0 2 0V10.63a2 2 0 0 1 2-2h8l-2.92 2.92a1 1 0 0 0 .71 1.71a1 1 0 0 0 .7-.3l4.63-4.62a1.19 1.19 0 0 0 .22-.34a1 1 0 0 0 0-.75Z"/>`),
		g.Group(children),
	)
}