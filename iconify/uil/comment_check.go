package uil

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func CommentCheck(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<path fill="currentColor" d="M13.29 7.13a1 1 0 0 0 0 1.42l1.92 1.92a1 1 0 0 0 1.42 0l4.08-4.08A1 1 0 1 0 19.29 5l-3.37 3.35l-1.21-1.22a1 1 0 0 0-1.42 0Zm6.62 3.51a1 1 0 0 0-.91 1.08a2.62 2.62 0 0 1 0 .28a7 7 0 0 1-7 7H6.41l.64-.63a1 1 0 0 0 0-1.41a7 7 0 0 1 3.48-11.81a7.14 7.14 0 0 1 2.8 0a1 1 0 1 0 .4-2a9.15 9.15 0 0 0-3.61 0A9.05 9.05 0 0 0 3 12a9 9 0 0 0 2 5.62l-1.71 1.67a1 1 0 0 0-.21 1.09A1 1 0 0 0 4 21h8a9 9 0 0 0 9-9v-.44a1 1 0 0 0-1.09-.92Z"/>`),
		g.Group(children),
	)
}