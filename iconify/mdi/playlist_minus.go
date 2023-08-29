package mdi

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func PlaylistMinus(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<path fill="currentColor" d="M3 16h7v-2H3m9 0v2h10v-2m-8-8H3v2h11m0 2H3v2h11v-2Z"/>`),
		g.Group(children),
	)
}