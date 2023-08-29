package mdi_light

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func Memory(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<path fill="currentColor" d="M8 19a3 3 0 0 1-3-3v-1H3v-1h2v-3H3v-1h2V9a3 3 0 0 1 3-3h1V4h1v2h3V4h1v2h1a3 3 0 0 1 3 3v1h2v1h-2v3h2v1h-2v1a3 3 0 0 1-3 3h-1v2h-1v-2h-3v2H9v-2H8M8 7a2 2 0 0 0-2 2v7a2 2 0 0 0 2 2h7a2 2 0 0 0 2-2V9a2 2 0 0 0-2-2H8m1 8v-5h5v5H9m1-4v3h3v-3h-3Z"/>`),
		g.Group(children),
	)
}