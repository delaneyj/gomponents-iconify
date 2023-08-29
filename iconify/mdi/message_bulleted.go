package mdi

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func MessageBulleted(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<path fill="currentColor" d="M20 2H4a2 2 0 0 0-2 2v18l4-4h14a2 2 0 0 0 2-2V4a2 2 0 0 0-2-2M8 14H6v-2h2v2m0-3H6V9h2v2m0-3H6V6h2v2m7 6h-5v-2h5v2m3-3h-8V9h8v2m0-3h-8V6h8v2Z"/>`),
		g.Group(children),
	)
}