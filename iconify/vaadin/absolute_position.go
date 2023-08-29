package vaadin

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func AbsolutePosition(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 0 0"),
		g.Raw(`<path fill="currentColor" d="M0 0v16h16V0H0zm15 15H1V9h3v1l3-2l-3-2v1H1V1h6v3H6l2 3l2-3H9V1h6v14z"/>`),
		g.Group(children),
	)
}