package tabler

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func SortDescendingNumbers(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<g fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="2"><path d="m4 15l3 3l3-3M7 6v12m10-4a2 2 0 0 1 2 2v3a2 2 0 1 1-4 0v-3a2 2 0 0 1 2-2zm-2-9a2 2 0 1 0 4 0a2 2 0 1 0-4 0"/><path d="M19 5v3a2 2 0 0 1-2 2h-1.5"/></g>`),
		g.Group(children),
	)
}