package iconoir

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func GitCompare(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<g fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="1.5"><path d="M18 21a2 2 0 1 0 0-4a2 2 0 0 0 0 4ZM6 7a2 2 0 1 0 0-4a2 2 0 0 0 0 4Zm12 10V7s0-2-2-2h-3M6 7v10s0 2 2 2h3"/><path d="M15 7.5L12.5 5L15 2.5m-6.5 14L11 19l-2.5 2.5"/></g>`),
		g.Group(children),
	)
}