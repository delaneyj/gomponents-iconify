package codicon

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func TriangleUp(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 0 0"),
		g.Raw(`<path fill="currentColor" d="m14 10.44l-.413.56H2.393L2 10.46L7.627 5h.827L14 10.44z"/>`),
		g.Group(children),
	)
}