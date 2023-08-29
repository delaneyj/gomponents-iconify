package mdi

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func PlaylistPlus(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<path fill="currentColor" d="M3 16h7v-2H3m15 0v-4h-2v4h-4v2h4v4h2v-4h4v-2m-8-8H3v2h11m0 2H3v2h11v-2Z"/>`),
		g.Group(children),
	)
}