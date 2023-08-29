package gg

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func GitFork(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<path fill="currentColor" d="M7 5a2 2 0 1 1 3 1.732v7.308h1.791a2 2 0 0 0 2-2v-1.256a2 2 0 1 1 2-.024v1.28a4 4 0 0 1-4 4H10v1.228A2 2 0 0 1 9 21a2 2 0 0 1-1-3.732V6.732A2 2 0 0 1 7 5Z"/>`),
		g.Group(children),
	)
}