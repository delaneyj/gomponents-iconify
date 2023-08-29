package vaadin

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func EnvelopeO(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 0 0"),
		g.Raw(`<path fill="currentColor" d="M0 3v11h16V3H0zm1 4.1l3.9 2L1 12.5V7.1zm.9 5.9l4-3.5L8 10.6l2.1-1.1l4 3.5H1.9zm13.1-.5L11.1 9L15 7v5.5zm0-6.6L8 9.4L1 5.9V4h14v1.9z"/>`),
		g.Group(children),
	)
}