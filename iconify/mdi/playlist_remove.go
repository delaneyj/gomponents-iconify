package mdi

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func PlaylistRemove(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<path fill="currentColor" d="M14 10H3v2h11v-2m0-4H3v2h11V6M3 16h7v-2H3v2m11.4 6l2.6-2.6l2.6 2.6l1.4-1.4l-2.6-2.6l2.6-2.6l-1.4-1.4l-2.6 2.6l-2.6-2.6l-1.4 1.4l2.6 2.6l-2.6 2.6l1.4 1.4Z"/>`),
		g.Group(children),
	)
}