package tabler

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func Tower(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<g fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="2"><path d="M5 3h1a1 1 0 0 1 1 1v2h3V4a1 1 0 0 1 1-1h2a1 1 0 0 1 1 1v2h3V4a1 1 0 0 1 1-1h1a1 1 0 0 1 1 1v4.394a2 2 0 0 1-.336 1.11l-1.328 1.992a2 2 0 0 0-.336 1.11V20a1 1 0 0 1-1 1H7a1 1 0 0 1-1-1v-7.394a2 2 0 0 0-.336-1.11L4.336 9.504A2 2 0 0 1 4 8.394V4a1 1 0 0 1 1-1z"/><path d="M10 21v-5a2 2 0 1 1 4 0v5"/></g>`),
		g.Group(children),
	)
}