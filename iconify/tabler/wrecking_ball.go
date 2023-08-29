package tabler

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func WreckingBall(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<g fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="2"><path d="M17 13a2 2 0 1 0 4 0a2 2 0 1 0-4 0M2 17a2 2 0 1 0 4 0a2 2 0 1 0-4 0m9 0a2 2 0 1 0 4 0a2 2 0 1 0-4 0m2 2H4m0-4h9"/><path d="M8 12V7h2a3 3 0 0 1 3 3v5"/><path d="M5 15v-2a1 1 0 0 1 1-1h7m6-1V4l-6 7"/></g>`),
		g.Group(children),
	)
}