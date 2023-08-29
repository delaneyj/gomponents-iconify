package iconamoon

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func NumberFourSquareLight(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<g fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="1.5"><path d="M4 4.001h16v14a2 2 0 0 1-2 2H6a2 2 0 0 1-2-2v-14Z"/><path d="m10.5 7l-1.272 4.45A2 2 0 0 0 11.152 14H14.5m0 0v-4m0 4v3"/></g>`),
		g.Group(children),
	)
}