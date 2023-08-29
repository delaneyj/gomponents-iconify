package mdi

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func CollapseAll(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<path fill="currentColor" d="M14 4H4v10H2V4a2 2 0 0 1 2-2h10v2m4 2H8a2 2 0 0 0-2 2v10h2V8h10V6m4 6v8a2 2 0 0 1-2 2h-8a2 2 0 0 1-2-2v-8a2 2 0 0 1 2-2h8a2 2 0 0 1 2 2m-2 3h-8v2h8v-2Z"/>`),
		g.Group(children),
	)
}