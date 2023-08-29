package ic

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func TwotoneSend(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<path fill="currentColor" d="m4 8.25l7.51 1l-7.5-3.22zm.01 9.72l7.5-3.22l-7.51 1z" opacity=".3"/><path fill="currentColor" d="M2.01 3L2 10l15 2l-15 2l.01 7L23 12L2.01 3zM4 8.25V6.03l7.51 3.22l-7.51-1zm.01 9.72v-2.22l7.51-1l-7.51 3.22z"/>`),
		g.Group(children),
	)
}