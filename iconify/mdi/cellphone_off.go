package mdi

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func CellphoneOff(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<path fill="currentColor" d="M2.38 1.73L1.11 3L5 6.89V21a2 2 0 0 0 2 2h10a2 2 0 0 0 2-2v-.11l1.84 1.84l1.27-1.27M17 19H7V8.89l10 10V19m0-14v8.8l2 2V3a2 2 0 0 0-2-2H7c-.72 0-1.4.37-1.76 1l3 3H17Z"/>`),
		g.Group(children),
	)
}