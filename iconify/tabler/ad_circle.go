package tabler

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func AdCircle(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<g fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="2"><path d="M2 12a10 10 0 1 0 20 0a10 10 0 1 0-20 0"/><path d="M7 15v-4.5a1.5 1.5 0 0 1 3 0V15m-3-2h3m4-4v6h1a2 2 0 0 0 2-2v-2a2 2 0 0 0-2-2h-1z"/></g>`),
		g.Group(children),
	)
}