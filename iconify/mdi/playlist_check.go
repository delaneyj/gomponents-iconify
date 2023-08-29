package mdi

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func PlaylistCheck(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<path fill="currentColor" d="M14 10H3v2h11v-2m0-4H3v2h11V6M3 16h7v-2H3v2m18.5-4.5L23 13l-7 7l-4.5-4.5L13 14l3 3l5.5-5.5Z"/>`),
		g.Group(children),
	)
}