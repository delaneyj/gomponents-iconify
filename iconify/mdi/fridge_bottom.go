package mdi

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func FridgeBottom(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<path fill="currentColor" d="M7 2a2 2 0 0 0-2 2v15a2 2 0 0 0 2 2v1h2v-1h6v1h2v-1a2 2 0 0 0 2-2V4a2 2 0 0 0-2-2H7m1 4h2v2H8V6m-1 5h10v8H7v-8m1 1v3h2v-3H8Z"/>`),
		g.Group(children),
	)
}