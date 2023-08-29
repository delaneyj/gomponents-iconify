package mdi

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func PacMan(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<path fill="currentColor" d="m12 12l7.07 7.07a9.953 9.953 0 0 1-14.14 0C1 15.17 1 8.84 4.93 4.93C8.83 1 15.16 1 19.07 4.93L12 12m7-2a2 2 0 0 0-2 2a2 2 0 0 0 2 2a2 2 0 0 0 2-2a2 2 0 0 0-2-2Z"/>`),
		g.Group(children),
	)
}