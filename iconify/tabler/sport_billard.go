package tabler

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func SportBillard(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<g fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="2"><path d="M10 10a2 2 0 1 0 4 0a2 2 0 1 0-4 0"/><path d="M10 14a2 2 0 1 0 4 0a2 2 0 1 0-4 0"/><path d="M4 12a8 8 0 1 0 16 0a8 8 0 1 0-16 0"/></g>`),
		g.Group(children),
	)
}