package mdi

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func TextSearchVariant(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<path fill="currentColor" d="M9 6v2H2V6h7m0 5v2H2v-2h7m9 5v2H2v-2h16m1.31-4.5c.44-.68.69-1.5.69-2.39c0-2.5-2-4.5-4.5-4.5s-4.5 2-4.5 4.5s2 4.5 4.5 4.5c.87 0 1.69-.25 2.38-.68L21 16l1.39-1.39l-3.08-3.11m-3.81.11c-1.38 0-2.5-1.11-2.5-2.5s1.12-2.5 2.5-2.5a2.5 2.5 0 0 1 0 5Z"/>`),
		g.Group(children),
	)
}