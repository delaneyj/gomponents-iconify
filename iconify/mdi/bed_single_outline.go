package mdi

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func BedSingleOutline(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<path fill="currentColor" d="M17 10V7a2 2 0 0 0-2-2H9c-1.1 0-2 .9-2 2v3c-1.1 0-2 .9-2 2v5h1.33L7 19h1l.67-2h6.66l.67 2h1l.67-2H19v-5a2 2 0 0 0-2-2M9 7h6v3H9m8 5H7v-3h10Z"/>`),
		g.Group(children),
	)
}