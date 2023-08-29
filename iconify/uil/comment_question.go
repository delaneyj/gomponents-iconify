package uil

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func CommentQuestion(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<path fill="currentColor" d="M14.77 5.87a1 1 0 0 0 1.36-.37A1 1 0 0 1 18 6a1 1 0 0 1-1 1a1 1 0 0 0 0 2a3 3 0 1 0-2.6-4.5a1 1 0 0 0 .37 1.37Zm4.3 7.13a1 1 0 0 0-1.12.86A7 7 0 0 1 11 20H5.41l.65-.65a1 1 0 0 0 0-1.41A7 7 0 0 1 11 6a1 1 0 0 0 0-2a9 9 0 0 0-7 14.61l-1.71 1.68a1 1 0 0 0-.21 1.09A1 1 0 0 0 3 22h8a9 9 0 0 0 8.93-7.87a1 1 0 0 0-.86-1.13Zm-1.69-2.93a1 1 0 0 0-.58-.07l-.18.06l-.18.09l-.15.13a1 1 0 0 0-.21.32a.84.84 0 0 0-.08.4a1 1 0 0 0 .07.39a1 1 0 0 0 .22.32A1 1 0 0 0 17 12a1 1 0 0 0 1-1a.84.84 0 0 0-.08-.38a1.07 1.07 0 0 0-.54-.54Z"/>`),
		g.Group(children),
	)
}