package tabler

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func WashMachine(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<g fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="2"><path d="M5 5a2 2 0 0 1 2-2h10a2 2 0 0 1 2 2v14a2 2 0 0 1-2 2H7a2 2 0 0 1-2-2z"/><path d="M8 14a4 4 0 1 0 8 0a4 4 0 1 0-8 0m0-8h.01M11 6h.01M14 6h2"/><path d="M8 14c1.333-.667 2.667-.667 4 0c1.333.667 2.667.667 4 0"/></g>`),
		g.Group(children),
	)
}