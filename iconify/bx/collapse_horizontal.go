package bx

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func CollapseHorizontal(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<path fill="currentColor" d="M19.95 5.64L13.59 12l6.36 6.36l1.41-1.41L16.41 12l4.95-4.95l-1.41-1.41zM2.64 7.05L7.59 12l-4.95 4.95l1.41 1.41L10.41 12L4.05 5.64L2.64 7.05z"/>`),
		g.Group(children),
	)
}