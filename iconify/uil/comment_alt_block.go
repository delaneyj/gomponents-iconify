package uil

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func CommentAltBlock(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<path fill="currentColor" d="M19 6h-4a1 1 0 0 0 0 2h4a1 1 0 0 1 1 1v9.72l-1.57-1.45a1 1 0 0 0-.68-.27H9a1 1 0 0 1-1-1v-1a1 1 0 0 0-2 0v1a3 3 0 0 0 3 3h8.36l3 2.73A1 1 0 0 0 21 22a1.1 1.1 0 0 0 .4-.08A1 1 0 0 0 22 21V9a3 3 0 0 0-3-3Zm-8.46 4.54A5 5 0 1 0 7 12a5 5 0 0 0 3.54-1.46ZM4 7a3 3 0 0 1 3-3a3 3 0 0 1 1.28.3l-4 4A3 3 0 0 1 4 7Zm5.7-1.29A3 3 0 0 1 10 7a3 3 0 0 1-4.28 2.7Z"/>`),
		g.Group(children),
	)
}