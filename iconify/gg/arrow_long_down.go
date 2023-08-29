package gg

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func ArrowLongDown(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<path fill="currentColor" d="m13.012 19.162l1.813-1.822l1.418 1.41l-4.231 4.255l-4.255-4.231l1.41-1.418l1.846 1.834L10.998.997l2-.002l.014 18.167Z"/>`),
		g.Group(children),
	)
}