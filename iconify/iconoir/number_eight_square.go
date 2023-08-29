package iconoir

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func NumberEightSquare(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<g fill="none" stroke="currentColor" stroke-width="1.5"><path d="M3 20.4V3.6a.6.6 0 0 1 .6-.6h16.8a.6.6 0 0 1 .6.6v16.8a.6.6 0 0 1-.6.6H3.6a.6.6 0 0 1-.6-.6Z"/><path stroke-linecap="round" stroke-linejoin="round" d="M12 16c-1.38 0-2.5-.5-2.5-2s1.12-2 2.5-2s2.5.5 2.5 2s-1.12 2-2.5 2Zm0-8c-1.38 0-2.5.5-2.5 2s1.12 2 2.5 2s2.5-.5 2.5-2s-1.12-2-2.5-2Z"/></g>`),
		g.Group(children),
	)
}