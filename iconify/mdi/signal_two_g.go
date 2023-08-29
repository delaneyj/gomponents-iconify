package mdi

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func SignalTwoG(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<path fill="currentColor" d="M11 19.5H2v-6a3 3 0 0 1 3-3h3v-3H2v-3h6a3 3 0 0 1 3 3v3a3 3 0 0 1-3 3H5v3h6m11-6h-4.5v3H19v3h-3v-9h6v-3h-6a3 3 0 0 0-3 3v9a3 3 0 0 0 3 3h3a3 3 0 0 0 3-3"/>`),
		g.Group(children),
	)
}