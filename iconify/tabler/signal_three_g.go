package tabler

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func SignalThreeG(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M17 8h-2a2 2 0 0 0-2 2v4a2 2 0 0 0 2 2h2v-4h-1M6 8h2.5A1.5 1.5 0 0 1 10 9.5v1A1.5 1.5 0 0 1 8.5 12H7h1.5a1.5 1.5 0 0 1 1.5 1.5v1A1.5 1.5 0 0 1 8.5 16H6"/>`),
		g.Group(children),
	)
}