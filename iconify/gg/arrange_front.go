package gg

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func ArrangeFront(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<path fill="currentColor" fill-rule="evenodd" d="M3 3h8v4h6v6h4v8h-8v-4H7v-6H3V3Zm12 6H9v6h6V9Z" clip-rule="evenodd"/>`),
		g.Group(children),
	)
}