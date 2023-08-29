package gg

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func DisplayFlex(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<g fill="currentColor"><path d="M6 17V7h2v10H6ZM16 7v10h2V7h-2Z"/><path fill-rule="evenodd" d="M2 3h20v18H2V3Zm2 2v14h16V5H4Z" clip-rule="evenodd"/></g>`),
		g.Group(children),
	)
}