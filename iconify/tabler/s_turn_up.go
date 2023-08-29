package tabler

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func STurnUp(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<g fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="2"><path d="M7 19a2 2 0 1 0-4 0a2 2 0 0 0 4 0z"/><path d="M5 17V7.5a3.5 3.5 0 0 1 7 0v9a3.5 3.5 0 0 0 7 0V3"/><path d="m16 6l3-3l3 3"/></g>`),
		g.Group(children),
	)
}