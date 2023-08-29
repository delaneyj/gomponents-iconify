package tabler

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func PlaneDeparture(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="m14.639 10.258l4.83-1.294a2 2 0 1 1 1.035 3.863L6.015 16.71l-4.45-5.02l2.897-.776l2.45 1.414l2.897-.776l-3.743-6.244l2.898-.777l5.675 5.727zM3 21h18"/>`),
		g.Group(children),
	)
}