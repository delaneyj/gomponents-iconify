package mdi

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func TriangleWave(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<path fill="currentColor" d="m22 12l-5 10L7.1 6.04L4.24 12H2L7 2l9.9 15.96L19.76 12H22Z"/>`),
		g.Group(children),
	)
}