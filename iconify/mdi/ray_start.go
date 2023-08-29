package mdi

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func RayStart(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<path fill="currentColor" d="M4 9c1.31 0 2.42.83 2.83 2H22v2H6.83A2.99 2.99 0 0 1 4 15a3 3 0 0 1-3-3a3 3 0 0 1 3-3Z"/>`),
		g.Group(children),
	)
}