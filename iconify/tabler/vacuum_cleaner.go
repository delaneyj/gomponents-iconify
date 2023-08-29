package tabler

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func VacuumCleaner(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<g fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="2"><path d="M21 12a9 9 0 1 1-18 0a9 9 0 0 1 18 0z"/><path d="M14 9a2 2 0 1 1-4 0a2 2 0 0 1 4 0zm-2 7h.01"/></g>`),
		g.Group(children),
	)
}