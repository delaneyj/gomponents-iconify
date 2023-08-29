package mdi

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func MessageOff(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<path fill="currentColor" d="M20.95 17.75L5.2 2H20a2 2 0 0 1 2 2v12c0 .76-.43 1.41-1.05 1.75M2.39 1.73L1.11 3l.89.9V22l4-4h10.11l4.73 4.73l1.27-1.27L2.39 1.73Z"/>`),
		g.Group(children),
	)
}