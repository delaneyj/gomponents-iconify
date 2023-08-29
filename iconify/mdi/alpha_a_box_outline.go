package mdi

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func AlphaABoxOutline(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<path fill="currentColor" d="M3 5a2 2 0 0 1 2-2h14a2 2 0 0 1 2 2v14a2 2 0 0 1-2 2H5a2 2 0 0 1-2-2V5m2 0v14h14V5H5m6 2h2a2 2 0 0 1 2 2v8h-2v-4h-2v4H9V9a2 2 0 0 1 2-2m0 2v2h2V9h-2Z"/>`),
		g.Group(children),
	)
}