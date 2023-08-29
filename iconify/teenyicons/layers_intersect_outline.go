package teenyicons

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func LayersIntersectOutline(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 15 15"),
		g.Raw(`<path fill="none" stroke="currentColor" d="M4.5 3V1.5a1 1 0 0 1 1-1H7m5 0h1.5a1 1 0 0 1 1 1V3M8 .5h3m1 10h1.5a1 1 0 0 0 1-1V8m0-4v3M3 4.5H1.5a1 1 0 0 0-1 1V7m0 5v1.5a1 1 0 0 0 1 1H3M.5 8v3M8 14.5h1.5a1 1 0 0 0 1-1V12M4 14.5h3m-2.5-10v6h6v-6h-6Z"/>`),
		g.Group(children),
	)
}