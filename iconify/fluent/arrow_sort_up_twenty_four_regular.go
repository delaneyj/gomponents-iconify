package fluent

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func ArrowSortUpTwentyFourRegular(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 20 20"),
		g.Raw(`<path fill="currentColor" d="m6.72 8.715l4.494-4.495a.75.75 0 0 1 .976-.073l.085.072l4.504 4.495a.75.75 0 0 1-.975 1.134l-.084-.072l-3.223-3.217v12.696a.75.75 0 0 1-.648.743l-.101.007a.75.75 0 0 1-.743-.648l-.007-.102l-.001-12.698L7.78 9.775a.75.75 0 0 1-.976.073l-.084-.073a.75.75 0 0 1-.073-.976l.073-.084l4.494-4.495L6.72 8.715Z"/>`),
		g.Group(children),
	)
}