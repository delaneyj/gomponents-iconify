package oi

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func Inbox(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 8 8"),
		g.Raw(`<path fill="currentColor" d="M.19 0C.08 0 0 .08 0 .19v7.63c0 .11.08.19.19.19h7.63c.11 0 .19-.08.19-.19V.19c0-.11-.08-.19-.19-.19H.19zM1 2h6v3H6L5 6H3L2 5H1V2z"/>`),
		g.Group(children),
	)
}