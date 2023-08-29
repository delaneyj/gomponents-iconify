package mdi

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func CircleMultiple(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<path fill="currentColor" d="M15 4a8 8 0 1 1-8 8a8 8 0 0 1 8-8M3 12a6 6 0 0 0 4 5.65v2.09A8 8 0 0 1 7 4.26v2.09A6 6 0 0 0 3 12Z"/>`),
		g.Group(children),
	)
}