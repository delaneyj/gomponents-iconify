package vaadin

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func PyramidChart(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 0 0"),
		g.Raw(`<path fill="currentColor" d="M10.29 5L8 1L5.71 5h4.58zm-8 6L0 15h16l-2.29-4H2.29zm10.85-1l-2.28-4H5.14l-2.28 4h10.28z"/>`),
		g.Group(children),
	)
}