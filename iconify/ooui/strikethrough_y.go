package ooui

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func StrikethroughY(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 20 20"),
		g.Raw(`<path fill="currentColor" d="M7.3 8h2.3l-4-6H3.3zm3.1 0h2.3l4-6h-2.3zM1 10v2h8v6h2v-6h8v-2z"/>`),
		g.Group(children),
	)
}