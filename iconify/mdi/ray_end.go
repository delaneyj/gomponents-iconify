package mdi

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func RayEnd(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<path fill="currentColor" d="M20 9c-1.31 0-2.42.83-2.83 2H2v2h15.17c.41 1.17 1.52 2 2.83 2a3 3 0 0 0 3-3a3 3 0 0 0-3-3Z"/>`),
		g.Group(children),
	)
}