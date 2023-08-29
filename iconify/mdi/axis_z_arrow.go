package mdi

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func AxisZArrow(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<path fill="currentColor" d="m12 2l4 4h-3v7.82l9.39 5.43l-1 1.75L12 15.56L2.61 21l-1-1.75L11 13.82V6H8l4-4Z"/>`),
		g.Group(children),
	)
}