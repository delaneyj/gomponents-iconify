package vaadin

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func AngleDoubleRight(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 0 0"),
		g.Raw(`<path fill="currentColor" d="M2 13h2l5-5l-5-5H2l5 5z"/><path fill="currentColor" d="M7 13h2l5-5l-5-5H7l5 5z"/>`),
		g.Group(children),
	)
}