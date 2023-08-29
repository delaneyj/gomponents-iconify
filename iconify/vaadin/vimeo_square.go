package vaadin

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func VimeoSquare(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 0 0"),
		g.Raw(`<path fill="currentColor" d="M0 0v16h16V0H0zm13.9 5.3c-.7 3.8-4.4 7-5.5 7.7s-2.2-.3-2.5-1.1c-.4-.9-1.7-5.7-2-6.1c-.4-.3-1.4.5-1.4.5L2 5.6s2-2.4 3.6-2.7s1.6 2.5 2 4.1c.4 1.5.6 2.4 1 2.4c.3 0 1-.9 1.7-2.2s0-2.5-1.4-1.6c.5-3.3 5.7-4.1 5-.3z"/>`),
		g.Group(children),
	)
}