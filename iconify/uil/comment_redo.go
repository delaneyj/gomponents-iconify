package uil

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func CommentRedo(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<path fill="currentColor" d="M19 14a1 1 0 0 0-1.22.72A7 7 0 0 1 11 20H5.41l.64-.63a1 1 0 0 0 0-1.41a7 7 0 0 1 3.2-11.74a1 1 0 0 0-.5-1.94A9 9 0 0 0 4 18.62l-1.71 1.67a1 1 0 0 0-.21 1.09A1 1 0 0 0 3 22h8a9 9 0 0 0 8.72-6.75A1 1 0 0 0 19 14Zm2-12a1 1 0 0 0-1 1a5 5 0 1 0 .3 7.75a1 1 0 1 0-1.3-1.5A3 3 0 1 1 17 4a3 3 0 0 1 2.23 1H18a1 1 0 0 0 0 2h3a1 1 0 0 0 1-1V3a1 1 0 0 0-1-1Z"/>`),
		g.Group(children),
	)
}