package bx

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func CommentX(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<path fill="currentColor" d="M9.707 13.707L12 11.414l2.293 2.293l1.414-1.414L13.414 10l2.293-2.293l-1.414-1.414L12 8.586L9.707 6.293L8.293 7.707L10.586 10l-2.293 2.293z"/><path fill="currentColor" d="M20 2H4c-1.103 0-2 .897-2 2v18l5.333-4H20c1.103 0 2-.897 2-2V4c0-1.103-.897-2-2-2zm0 14H6.667L4 18V4h16v12z"/>`),
		g.Group(children),
	)
}