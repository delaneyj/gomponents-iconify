package mdi

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func AlphaABox(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<path fill="currentColor" d="M3 5a2 2 0 0 1 2-2h14a2 2 0 0 1 2 2v14a2 2 0 0 1-2 2H5a2 2 0 0 1-2-2V5m8 2a2 2 0 0 0-2 2v8h2v-4h2v4h2V9a2 2 0 0 0-2-2h-2m0 2h2v2h-2V9Z"/>`),
		g.Group(children),
	)
}