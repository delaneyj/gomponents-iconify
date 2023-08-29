package mdi

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func ApplicationSettings(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<path fill="currentColor" d="M21 0H3C1.9 0 1 .9 1 2v16c0 1.1.9 2 2 2h18c1.1 0 2-.9 2-2V2c0-1.1-.9-2-2-2m0 5H3V2h18v3M7 22h2v2H7v-2m4 0h2v2h-2v-2m4 0h2v2h-2v-2"/>`),
		g.Group(children),
	)
}