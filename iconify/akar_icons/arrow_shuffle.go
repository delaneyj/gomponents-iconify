package akar_icons

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func ArrowShuffle(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<g fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="2"><path d="M2 19h3.908a2 2 0 0 0 1.682-.919L11.5 12l3.91-6.082A2 2 0 0 1 17.091 5H22m0 14h-4.908a2 2 0 0 1-1.682-.919L13.428 15M2 5h3.908a2 2 0 0 1 1.682.918L9.571 9"/><path d="m19 2l3 3l-3 3m0 8l3 3l-3 3"/></g>`),
		g.Group(children),
	)
}