package gg

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func ArrangeBack(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<path fill="currentColor" fill-rule="evenodd" d="M3 3h8v4h6v6h4v8h-8v-4H7v-6H3V3Zm12 10h-2v2H9v-4h2V9h4v4Z" clip-rule="evenodd"/>`),
		g.Group(children),
	)
}