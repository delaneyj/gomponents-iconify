package vaadin

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func Mailbox(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 0 0"),
		g.Raw(`<path fill="currentColor" d="M13 1H3l-3 9v5h16v-5l-3-9zm-2 9v2H5v-2H1.1l2.7-8h8.6l2.7 8H11z"/>`),
		g.Group(children),
	)
}