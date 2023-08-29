package bx

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func CollapseVertical(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<path fill="currentColor" d="M12 7.59L7.05 2.64L5.64 4.05L12 10.41l6.36-6.36l-1.41-1.41L12 7.59zM5.64 19.95l1.41 1.41L12 16.41l4.95 4.95l1.41-1.41L12 13.59l-6.36 6.36z"/>`),
		g.Group(children),
	)
}