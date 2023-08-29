package mdi

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func BedKingOutline(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<path fill="currentColor" d="M20 10V7a2 2 0 0 0-2-2H6a2 2 0 0 0-2 2v3a2 2 0 0 0-2 2v5h1.33L4 19h1l.67-2h12.66l.67 2h1l.67-2H22v-5a2 2 0 0 0-2-2m-7-3h5v3h-5M6 7h5v3H6m14 5H4v-3h16Z"/>`),
		g.Group(children),
	)
}