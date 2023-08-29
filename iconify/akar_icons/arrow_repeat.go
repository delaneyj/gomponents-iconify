package akar_icons

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func ArrowRepeat(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<g fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="2"><path d="m18 2l3 3l-3 3M6 22l-3-3l3-3"/><path d="M21 5H10a7 7 0 0 0-7 7m0 7h11a7 7 0 0 0 7-7"/></g>`),
		g.Group(children),
	)
}