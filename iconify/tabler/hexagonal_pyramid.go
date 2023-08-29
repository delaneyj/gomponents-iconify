package tabler

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func HexagonalPyramid(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M11.162 2.457L3.316 15.411a1.988 1.988 0 0 0 .267 2.483l2.527 2.523c.374.373.88.583 1.408.583h8.964c.528 0 1.034-.21 1.408-.583l2.527-2.523a1.988 1.988 0 0 0 .267-2.483L12.838 2.457a.996.996 0 0 0-1.676 0zM12 2L7 20.9M12 2l5 18.9"/>`),
		g.Group(children),
	)
}