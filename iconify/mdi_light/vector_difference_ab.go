package mdi_light

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func VectorDifferenceAb(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<path fill="currentColor" d="M18 8a3 3 0 0 1 3 3v8a3 3 0 0 1-3 3h-8a3 3 0 0 1-3-3v-4h5a2 2 0 0 0 2-2V8h4M8 16v3a2 2 0 0 0 2 2h8a2 2 0 0 0 2-2v-8a2 2 0 0 0-2-2h-3v4a3 3 0 0 1-3 3H8m3-14h1a3 3 0 0 1 3 3v1h-1V5a2 2 0 0 0-2-2h-1V2M4 2h1v1H4a2 2 0 0 0-2 2v1H1V5a3 3 0 0 1 3-3m3 1V2h2v1H7m-5 7H1V8h1v2m0 3a2 2 0 0 0 2 2h1v1H4a3 3 0 0 1-3-3v-1h1v1Z"/>`),
		g.Group(children),
	)
}