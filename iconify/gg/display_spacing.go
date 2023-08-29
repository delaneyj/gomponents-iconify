package gg

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func DisplaySpacing(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<g fill="currentColor"><path d="M3 21V3h2v18H3Z"/><path fill-rule="evenodd" d="M7 3h10v18H7V3Zm2 2v14h6V5H9Z" clip-rule="evenodd"/><path d="M19 3v18h2V3h-2Z"/></g>`),
		g.Group(children),
	)
}