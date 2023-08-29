package majesticons

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func ViewRowsLine(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M3 10v6a2 2 0 0 0 2 2h14a2 2 0 0 0 2-2v-6M3 10V8a2 2 0 0 1 2-2h14a2 2 0 0 1 2 2v2M3 10h18M3 14h18"/>`),
		g.Group(children),
	)
}