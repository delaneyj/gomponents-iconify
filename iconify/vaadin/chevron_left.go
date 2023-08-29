package vaadin

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func ChevronLeft(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 0 0"),
		g.Raw(`<path fill="currentColor" d="m2.9 8l8-8l2.2 2.1L7.2 8l5.9 5.9l-2.2 2.1z"/>`),
		g.Group(children),
	)
}