package uil

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func CommentAltRedo(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<path fill="currentColor" d="M10.3 10.75A1 1 0 1 0 9 9.25A3 3 0 1 1 7 4a3 3 0 0 1 2.23 1H8a1 1 0 0 0 0 2h3a1 1 0 0 0 1-1V3a1 1 0 0 0-2 0a5 5 0 1 0 .3 7.75ZM19 6h-4a1 1 0 0 0 0 2h4a1 1 0 0 1 1 1v9.72l-1.57-1.45a1 1 0 0 0-.68-.27H9a1 1 0 0 1-1-1v-1a1 1 0 0 0-2 0v1a3 3 0 0 0 3 3h8.36l3 2.73A1 1 0 0 0 21 22a1.1 1.1 0 0 0 .4-.08A1 1 0 0 0 22 21V9a3 3 0 0 0-3-3Z"/>`),
		g.Group(children),
	)
}