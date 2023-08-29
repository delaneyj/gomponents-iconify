package mdi

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func PlaylistPlay(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<path fill="currentColor" d="M3 10h11v2H3v-2m0-4h11v2H3V6m0 8h7v2H3v-2m13-1v8l6-4l-6-4Z"/>`),
		g.Group(children),
	)
}