package bx

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func ArrowFromBottom(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<path fill="currentColor" d="M6 18h12v2H6zm6-14.414l-6.707 6.707l1.414 1.414L11 7.414V16h2V7.414l4.293 4.293l1.414-1.414z"/>`),
		g.Group(children),
	)
}