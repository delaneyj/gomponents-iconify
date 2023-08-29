package uil

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func CommentAltQuestion(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<path fill="currentColor" d="M6.29 10.3a1 1 0 0 0 1.09 1.63a1.19 1.19 0 0 0 .33-.22a1 1 0 0 0 .21-.32A.85.85 0 0 0 8 11a1 1 0 0 0-.29-.7a1 1 0 0 0-1.42 0ZM7 5a1 1 0 0 1 0 2a1 1 0 0 0 0 2a3 3 0 1 0-2.6-4.5a1 1 0 0 0 .37 1.37a1 1 0 0 0 1.36-.37A1 1 0 0 1 7 5Zm12 1h-6a1 1 0 0 0 0 2h6a1 1 0 0 1 1 1v9.72l-1.57-1.45a1 1 0 0 0-.68-.27H9a1 1 0 0 1-1-1v-1a1 1 0 0 0-2 0v1a3 3 0 0 0 3 3h8.36l3 2.73A1 1 0 0 0 21 22a1.1 1.1 0 0 0 .4-.08A1 1 0 0 0 22 21V9a3 3 0 0 0-3-3Z"/>`),
		g.Group(children),
	)
}