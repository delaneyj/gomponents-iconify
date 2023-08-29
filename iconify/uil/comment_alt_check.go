package uil

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func CommentAltCheck(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<path fill="currentColor" d="m6.77 9.15l5.44-5.44a1 1 0 1 0-1.42-1.42L6.06 7L4.21 5.17a1 1 0 0 0-1.42 1.42l2.56 2.56a1 1 0 0 0 1.42 0ZM18.5 6H13a1 1 0 0 0 0 2h5.5a1 1 0 0 1 1 1v9.72l-1.57-1.45a1 1 0 0 0-.68-.27H8.5a1 1 0 0 1-1-1v-3.5a1 1 0 0 0-2 0V16a3 3 0 0 0 3 3h8.36l3 2.73a1 1 0 0 0 .68.27a1.1 1.1 0 0 0 .4-.08a1 1 0 0 0 .6-.92V9a3 3 0 0 0-3.04-3Z"/>`),
		g.Group(children),
	)
}