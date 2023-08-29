package vaadin

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func ChevronRight(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 0 0"),
		g.Raw(`<path fill="currentColor" d="m13.1 8l-8 8l-2.2-2.1L8.8 8L2.9 2.1L5.1 0z"/>`),
		g.Group(children),
	)
}