package iconamoon

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func NumberZeroSquare(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<g fill="none" stroke="currentColor" stroke-linejoin="round" stroke-width="2"><path stroke-linecap="round" d="M4 4.001h16v14a2 2 0 0 1-2 2H6a2 2 0 0 1-2-2v-14Z"/><path d="M12 17c1.657 0 3-1.4 3-3.125v-3.75C15 8.399 13.657 7 12 7s-3 1.4-3 3.125v3.75C9 15.601 10.343 17 12 17Z"/></g>`),
		g.Group(children),
	)
}