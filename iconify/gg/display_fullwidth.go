package gg

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func DisplayFullwidth(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<g fill="currentColor"><path d="M2 5h20V3H2v2Zm0 16h20v-2H2v2Z"/><path fill-rule="evenodd" d="M2 7v10h20V7H2Zm2 2h16v6H4V9Z" clip-rule="evenodd"/></g>`),
		g.Group(children),
	)
}