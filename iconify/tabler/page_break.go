package tabler

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func PageBreak(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<g fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="2"><path d="M14 3v4a1 1 0 0 0 1 1h4m0 10v1a2 2 0 0 1-2 2H7a2 2 0 0 1-2-2v-1m-2-4h3m4.5 0h3m4.5 0h3"/><path d="M5 10V5a2 2 0 0 1 2-2h7l5 5v2"/></g>`),
		g.Group(children),
	)
}