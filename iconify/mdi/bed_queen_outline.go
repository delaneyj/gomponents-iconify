package mdi

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func BedQueenOutline(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<path fill="currentColor" d="M19 10V7a2 2 0 0 0-2-2H7a2 2 0 0 0-2 2v3a2 2 0 0 0-2 2v5h1.33L5 19h1l.67-2h10.66l.67 2h1l.67-2H21v-5a2 2 0 0 0-2-2M7 7h10v3H7m12 5H5v-3h14Z"/>`),
		g.Group(children),
	)
}