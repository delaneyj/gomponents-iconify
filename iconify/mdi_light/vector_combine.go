package mdi_light

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func VectorCombine(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<path fill="currentColor" d="M4 2h8a3 3 0 0 1 3 3v3h3a3 3 0 0 1 3 3v8a3 3 0 0 1-3 3h-8a3 3 0 0 1-3-3v-3H4a3 3 0 0 1-3-3V5a3 3 0 0 1 3-3m11 11a3 3 0 0 1-3 3H8v3a2 2 0 0 0 2 2h8a2 2 0 0 0 2-2v-8a2 2 0 0 0-2-2h-3v4M4 3a2 2 0 0 0-2 2v8a2 2 0 0 0 2 2h3v-4a3 3 0 0 1 3-3h4V5a2 2 0 0 0-2-2H4m8 12a2 2 0 0 0 2-2V9h-4a2 2 0 0 0-2 2v4h4Z"/>`),
		g.Group(children),
	)
}