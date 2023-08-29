package tabler

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func SortAscendingLetters(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M15 10V5c0-1.38.62-2 2-2s2 .62 2 2v5m0-3h-4m4 14h-4l4-7h-4M4 15l3 3l3-3M7 6v12"/>`),
		g.Group(children),
	)
}