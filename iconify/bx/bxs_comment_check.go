package bx

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func BxsCommentCheck(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<path d="M20 2H4c-1.103 0-2 .897-2 2v18l4-4h14c1.103 0 2-.897 2-2V4c0-1.103-.897-2-2-2zm-9 11.914l-3.707-3.707l1.414-1.414L11 11.086l4.793-4.793l1.414 1.414L11 13.914z" fill="currentColor"/>`),
		g.Group(children),
	)
}