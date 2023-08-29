package gg

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func Cast(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<g fill="currentColor"><path d="M20 6H4v2H2V6a2 2 0 0 1 2-2h16a2 2 0 0 1 2 2v12a2 2 0 0 1-2 2h-5v-2h5V6ZM2 13a7 7 0 0 1 7 7H7a5 5 0 0 0-5-5v-2Zm0 4a3 3 0 0 1 3 3H2v-3Z"/><path d="M2 9c6.075 0 11 4.925 11 11h-2a9 9 0 0 0-9-9V9Z"/></g>`),
		g.Group(children),
	)
}