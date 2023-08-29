package tabler

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func HexagonalPyramidPlus(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="m18.642 12.04l-5.804-9.583a.996.996 0 0 0-1.676 0L3.316 15.411a1.988 1.988 0 0 0 .267 2.483l2.527 2.523c.374.373.88.583 1.408.583H12.5M12 2L7 20.9M12 2l3.304 12.489M16 19h6m-3-3v6"/>`),
		g.Group(children),
	)
}