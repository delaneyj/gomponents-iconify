package gg

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func ShapeTriangle(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<path fill="currentColor" fill-rule="evenodd" d="M11.66 5L3 20h17.32L11.66 5Zm0 6l-3.464 6h6.928l-3.464-6Z" clip-rule="evenodd"/>`),
		g.Group(children),
	)
}