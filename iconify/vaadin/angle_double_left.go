package vaadin

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func AngleDoubleLeft(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 0 0"),
		g.Raw(`<path fill="currentColor" d="M14 3h-2L7 8l5 5h2L9 8z"/><path fill="currentColor" d="M9 3H7L2 8l5 5h2L4 8z"/>`),
		g.Group(children),
	)
}