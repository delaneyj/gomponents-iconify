package iconoir

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func Calculator(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<g fill="none" stroke="currentColor" stroke-width="1.5"><path d="M1 21V3a2 2 0 0 1 2-2h18a2 2 0 0 1 2 2v18a2 2 0 0 1-2 2H3a2 2 0 0 1-2-2Z"/><path stroke-linecap="round" stroke-linejoin="round" d="M15 7h4m-4 8.5h4m-4 3h4M5 7h2m2 0H7m0 0V5m0 2v2m-1.414 9.414L7 17m1.415-1.414L7 17m0 0l-1.414-1.414M7 17l1.415 1.414"/></g>`),
		g.Group(children),
	)
}