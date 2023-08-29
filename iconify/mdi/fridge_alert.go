package mdi

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func FridgeAlert(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<path fill="currentColor" d="M5 2h10a2 2 0 0 1 2 2v5H3V4c0-1.1.9-2 2-2m12 17c0 1.11-.89 2-2 2v1h-2v-1H7v1H5v-1a2 2 0 0 1-2-2v-9h14v9M6 5v2h2V5H6m0 7v3h2v-3H6m13 3h2v2h-2v-2m0-8h2v6h-2V7Z"/>`),
		g.Group(children),
	)
}