package bx

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func BxsCommentX(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<path d="M20 2H4c-1.103 0-2 .897-2 2v18l4-4h14c1.103 0 2-.897 2-2V4c0-1.103-.897-2-2-2zm-3.294 11.543l-1.414 1.414l-3.293-3.292l-3.292 3.292l-1.414-1.414l3.292-3.292l-3.292-3.293l1.414-1.414l3.292 3.292l3.293-3.292l1.414 1.414l-3.292 3.293l3.292 3.292z" fill="currentColor"/>`),
		g.Group(children),
	)
}