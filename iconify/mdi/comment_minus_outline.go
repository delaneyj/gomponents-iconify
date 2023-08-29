package mdi

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func CommentMinusOutline(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<path fill="currentColor" d="M16 11H8V9h8v2m6-7v12c0 1.11-.89 2-2 2h-6.1l-3.7 3.71c-.2.19-.45.29-.7.29H9c-.55 0-1-.45-1-1v-3H4a2 2 0 0 1-2-2V4a2 2 0 0 1 2-2h16a2 2 0 0 1 2 2m-2 0H4v12h6v3.08L13.08 16H20V4"/>`),
		g.Group(children),
	)
}