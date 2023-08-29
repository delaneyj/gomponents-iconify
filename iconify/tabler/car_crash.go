package tabler

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func CarCrash(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<g fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="2"><path d="M8 17a2 2 0 1 0 4 0a2 2 0 1 0-4 0"/><path d="m7 6l4 5h1a2 2 0 0 1 2 2v4h-2m-4 0H3m0-6h8m-6 0V6m2 0H3m11 2V6m5 6h2m-3.5 3.5L19 17m-1.5-8.5L19 7"/></g>`),
		g.Group(children),
	)
}