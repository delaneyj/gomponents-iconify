package mdi

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func HeadsetDock(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<path fill="currentColor" d="M2 18h7V6.13C7.27 6.57 6 8.14 6 10v1h2v6H6a2 2 0 0 1-2-2v-5a6 6 0 0 1 6-6h1a6 6 0 0 1 6 6v2h1V9h2v3a2 2 0 0 1-2 2h-1v1a2 2 0 0 1-2 2h-2v-6h2v-1c0-1.86-1.27-3.43-3-3.87V18h10v2H2v-2Z"/>`),
		g.Group(children),
	)
}