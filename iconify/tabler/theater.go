package tabler

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func Theater(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M4 20h16m0-4V6a2 2 0 0 0-2-2H6a2 2 0 0 0-2 2v10l4-6c2.667 1.333 5.333 1.333 8 0l4 6z"/>`),
		g.Group(children),
	)
}