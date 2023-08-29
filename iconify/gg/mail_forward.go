package gg

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func MailForward(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<path fill="currentColor" d="m13.638 15.529l1.414 1.414l4.95-4.95l-4.95-4.95l-1.414 1.415l2.498 2.498H7.998a4 4 0 0 0-4 4v2h2v-2a2 2 0 0 1 2-2h8.212l-2.572 2.573Z"/>`),
		g.Group(children),
	)
}