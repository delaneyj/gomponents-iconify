package gg

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func PathIntersect(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<path fill="currentColor" fill-rule="evenodd" d="M15 5H5v10h4v4h10V9h-4V5Zm-2 2H7v6h2V9h4V7Zm4 10h-6v-2h4v-4h2v6Z" clip-rule="evenodd"/>`),
		g.Group(children),
	)
}