package tabler

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func NoCreativeCommons(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<g fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="2"><path d="M3 12a9 9 0 1 0 18 0a9 9 0 1 0-18 0"/><path d="M10.5 10.5a2.187 2.187 0 0 0-2.914.116a1.928 1.928 0 0 0 0 2.768a2.188 2.188 0 0 0 2.914.116m6-3a2.187 2.187 0 0 0-2.914.116a1.928 1.928 0 0 0 0 2.768a2.188 2.188 0 0 0 2.914.116M6 6l1.5 1.5m9 9L18 18"/></g>`),
		g.Group(children),
	)
}