package iconamoon

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func SignYBold(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<g fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="2.5"><path d="M7 21h.525a2.64 2.64 0 0 0 2.248-1.256l6.29-10.22A1 1 0 0 0 15.21 8H14"/><path d="M12 16L9.487 9.298A2 2 0 0 0 7.614 8H7"/></g>`),
		g.Group(children),
	)
}