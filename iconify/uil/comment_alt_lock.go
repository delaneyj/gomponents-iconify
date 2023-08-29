package uil

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func CommentAltLock(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<path fill="currentColor" d="M12 10V8a3 3 0 0 0-2-2.82V4a3 3 0 0 0-6 0v1.18A3 3 0 0 0 2 8v2a3 3 0 0 0 3 3h4a3 3 0 0 0 3-3ZM6 4a1 1 0 0 1 2 0v1H6Zm-2 6V8a1 1 0 0 1 1-1h4a1 1 0 0 1 1 1v2a1 1 0 0 1-1 1H5a1 1 0 0 1-1-1Zm15-3h-4a1 1 0 0 0 0 2h4a1 1 0 0 1 1 1v9.72l-1.57-1.45a1 1 0 0 0-.68-.27H9a1 1 0 0 1-1-1v-1a1 1 0 0 0-2 0v1a3 3 0 0 0 3 3h8.36l3 2.73A1 1 0 0 0 21 23a1.1 1.1 0 0 0 .4-.08A1 1 0 0 0 22 22V10a3 3 0 0 0-3-3Z"/>`),
		g.Group(children),
	)
}