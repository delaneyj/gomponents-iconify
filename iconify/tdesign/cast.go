package tdesign

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func Cast(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<path fill="currentColor" d="M1 3h22v18H13v-2h8V5H3v4H1V3Zm0 8h1a9 9 0 0 1 9 9v1H9v-1a7 7 0 0 0-7-7H1v-2Zm0 4h1a5 5 0 0 1 5 5v1H5v-1a3 3 0 0 0-3-3H1v-2Zm0 4h2.004v2.004H1V19Z"/>`),
		g.Group(children),
	)
}