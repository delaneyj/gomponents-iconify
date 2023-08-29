package gg

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func ListTree(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<path fill="currentColor" fill-rule="evenodd" d="M9 1H1v8h8V6h2v14h4v3h8v-8h-8v3h-2V6h2v3h8V1h-8v3H9V1Zm12 2h-4v4h4V3Zm-4 14h4v4h-4v-4Z" clip-rule="evenodd"/>`),
		g.Group(children),
	)
}