package uil

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func CommentSlash(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<path fill="currentColor" d="M10.19 5.23A7.12 7.12 0 0 1 12 5a7 7 0 0 1 7 7a7.12 7.12 0 0 1-.23 1.81a1 1 0 0 0 .7 1.23a1.15 1.15 0 0 0 .26 0a1 1 0 0 0 1-.74A8.91 8.91 0 0 0 21 12a9 9 0 0 0-9-9a8.91 8.91 0 0 0-2.33.3A1 1 0 0 0 9 4.53a1 1 0 0 0 1.19.7Zm11.52 15.06l-18-18a1 1 0 0 0-1.42 1.42L5 6.38a9 9 0 0 0 0 11.24l-1.71 1.67a1 1 0 0 0-.21 1.09A1 1 0 0 0 4 21h8a9 9 0 0 0 5.62-2l2.67 2.68a1 1 0 0 0 1.42 0a1 1 0 0 0 0-1.39ZM12 19H6.41l.64-.63a1 1 0 0 0 0-1.41a7 7 0 0 1-.65-9.15l9.79 9.79A7 7 0 0 1 12 19Z"/>`),
		g.Group(children),
	)
}