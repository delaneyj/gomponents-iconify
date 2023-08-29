package gg

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func MailReply(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<path fill="currentColor" d="m10.362 15.529l-1.414 1.414l-4.95-4.95l4.95-4.95l1.414 1.415l-2.498 2.498h8.138a4 4 0 0 1 4 4v2h-2v-2a2 2 0 0 0-2-2H7.79l2.572 2.573Z"/>`),
		g.Group(children),
	)
}