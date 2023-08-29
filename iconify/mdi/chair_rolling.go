package mdi

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func ChairRolling(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<path fill="currentColor" d="M22 10v3h-3v-3h3M2 13h3v-3H2v3m15-8c0-1.1-.9-2-2-2H9c-1.1 0-2 .9-2 2v8h10V5M7 15H6v2h5v1l-4 4h2.8l2.2-2.2l2.2 2.2H17l-4-4v-1h5v-2H7Z"/>`),
		g.Group(children),
	)
}