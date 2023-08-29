package tabler

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func Asset(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<g fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="2"><path d="M3 15a6 6 0 1 0 12 0a6 6 0 1 0-12 0"/><path d="M7 15a2 2 0 1 0 4 0a2 2 0 1 0-4 0M17 5a2 2 0 1 0 4 0a2 2 0 1 0-4 0m-2.782 12.975l6.619-12.174M6.079 9.756l12.217-6.631"/><path d="M7 15a2 2 0 1 0 4 0a2 2 0 1 0-4 0"/></g>`),
		g.Group(children),
	)
}