package gg

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func ArrowTopLeftO(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<g fill="currentColor"><path d="M10 13.963H8v-6h6v2h-2.586l5.33 5.33l-1.414 1.414l-5.33-5.33v2.586Z"/><path fill-rule="evenodd" d="M23 12c0-6.075-4.925-11-11-11S1 5.925 1 12s4.925 11 11 11s11-4.925 11-11Zm-2 0a9 9 0 1 0-18 0a9 9 0 0 0 18 0Z" clip-rule="evenodd"/></g>`),
		g.Group(children),
	)
}