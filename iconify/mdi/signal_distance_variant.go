package mdi

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func SignalDistanceVariant(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<path fill="currentColor" d="M4 6V4a12 12 0 0 1 12 12h-2A10 10 0 0 0 4 6m0 4V8a8 8 0 0 1 8 8h-2a6 6 0 0 0-6-6m0 2a4 4 0 0 1 4 4H4v-4m-1 6h16v-2l3 3l-3 3v-2H3v-2Z"/>`),
		g.Group(children),
	)
}