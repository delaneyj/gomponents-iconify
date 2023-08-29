package humbleicons

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func Flask(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<g fill="none" stroke="currentColor" stroke-width="2"><path stroke-linejoin="round" d="M9.5 10V4h5v6l4.743 5.174A2.88 2.88 0 0 1 17.12 20H6.88a2.88 2.88 0 0 1-2.123-4.826L9.5 10z"/><path stroke-linecap="round" d="M8.5 4h7"/><path d="M6 14c3.5 2 4 0 6 0s2.5 2 6 0"/></g>`),
		g.Group(children),
	)
}