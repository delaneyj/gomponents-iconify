package tabler

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func Tournament(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<g fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="2"><path d="M2 4a2 2 0 1 0 4 0a2 2 0 1 0-4 0m16 6a2 2 0 1 0 4 0a2 2 0 1 0-4 0M2 12a2 2 0 1 0 4 0a2 2 0 1 0-4 0m0 8a2 2 0 1 0 4 0a2 2 0 1 0-4 0m4-8h3a1 1 0 0 1 1 1v6a1 1 0 0 1-1 1H6"/><path d="M6 4h7a1 1 0 0 1 1 1v10a1 1 0 0 1-1 1h-2m3-6h4"/></g>`),
		g.Group(children),
	)
}