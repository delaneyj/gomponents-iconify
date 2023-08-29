package tabler

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func STurnLeft(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<g fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="2"><path d="M19 7a2 2 0 1 1 0-4a2 2 0 0 1 0 4z"/><path d="M17 5H7.5a3.5 3.5 0 0 0 0 7h9a3.5 3.5 0 0 1 0 7H3"/><path d="m6 16l-3 3l3 3"/></g>`),
		g.Group(children),
	)
}