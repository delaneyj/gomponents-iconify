package ic

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func OutlineSend(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<path fill="currentColor" d="m4.01 6.03l7.51 3.22l-7.52-1l.01-2.22m7.5 8.72L4 17.97v-2.22l7.51-1M2.01 3L2 10l15 2l-15 2l.01 7L23 12L2.01 3z"/>`),
		g.Group(children),
	)
}