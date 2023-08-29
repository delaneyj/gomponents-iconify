package tabler

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func SquareFThree(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<g fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="2"><path d="M3 5a2 2 0 0 1 2-2h14a2 2 0 0 1 2 2v14a2 2 0 0 1-2 2H5a2 2 0 0 1-2-2z"/><path d="M13 9.5a.5.5 0 0 1 .5-.5h1a1.5 1.5 0 0 1 0 3H14h.5a1.5 1.5 0 0 1 0 3h-1a.5.5 0 0 1-.5-.5M8 12h2m0-3H8v6"/></g>`),
		g.Group(children),
	)
}