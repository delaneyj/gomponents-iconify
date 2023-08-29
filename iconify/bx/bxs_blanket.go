package bx

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func BxsBlanket(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<path d="M20 2H7C4.243 2 2 4.243 2 7v10c0 2.757 2.243 5 5 5h12c1.654 0 3-1.346 3-3s-1.346-3-3-3H6v2h13a1 1 0 0 1 0 2H7c-1.654 0-3-1.346-3-3s1.346-3 3-3h13c1.103 0 2-.897 2-2V4c0-1.103-.897-2-2-2zM4 13h.003L4 13.002V13z" fill="currentColor"/>`),
		g.Group(children),
	)
}