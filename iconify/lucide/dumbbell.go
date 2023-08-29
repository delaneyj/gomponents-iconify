package lucide

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func Dumbbell(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="m6.5 6.5l11 11M21 21l-1-1M3 3l1 1m14 18l4-4M2 6l4-4m-3 8l7-7m4 18l7-7"/>`),
		g.Group(children),
	)
}