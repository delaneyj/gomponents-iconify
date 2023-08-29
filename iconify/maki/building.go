package maki

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func Building(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 15 15"),
		g.Raw(`<path fill="currentColor" d="M3 2v11h5v-3h3v3h1V2H3zm4 10H4v-2h3v2zm0-3H4V7h3v2zm0-3H4V4h3v2zm4 3H8V7h3v2zm0-3H8V4h3v2z"/>`),
		g.Group(children),
	)
}