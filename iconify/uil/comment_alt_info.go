package uil

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func CommentAltInfo(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<path fill="currentColor" d="M5 11a1 1 0 0 0 1-1V6a1 1 0 0 0-2 0v4a1 1 0 0 0 1 1Zm-.71-7.29a1 1 0 0 0 1.09.21a1.15 1.15 0 0 0 .33-.21a1.15 1.15 0 0 0 .21-.33A.84.84 0 0 0 6 3a1 1 0 0 0-.29-.71a1 1 0 0 0-1.09-.21a1 1 0 0 0-.33.21A1 1 0 0 0 4 3a1 1 0 0 0 .08.38a1.15 1.15 0 0 0 .21.33ZM17 6H9a1 1 0 0 0 0 2h8a1 1 0 0 1 1 1v9.72l-1.57-1.45a1 1 0 0 0-.68-.27H7a1 1 0 0 1-1-1v-2a1 1 0 0 0-2 0v2a3 3 0 0 0 3 3h8.36l3 2.73A1 1 0 0 0 19 22a1.1 1.1 0 0 0 .4-.08A1 1 0 0 0 20 21V9a3 3 0 0 0-3-3Z"/>`),
		g.Group(children),
	)
}