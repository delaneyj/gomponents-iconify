package tabler

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func Horse(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="m7 10l-.85 8.507A1.357 1.357 0 0 0 7.5 20h.146a2 2 0 0 0 1.857-1.257l.994-2.486A2 2 0 0 1 12.354 15h1.292a2 2 0 0 1 1.857 1.257l.994 2.486A2 2 0 0 0 18.354 20h.146a1.37 1.37 0 0 0 1.364-1.494L19 9h-8c0-3-3-5-6-5l-3 6l2 2l3-2zm15 4v-2a3 3 0 0 0-3-3"/>`),
		g.Group(children),
	)
}