package majesticons

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func ServerLine(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M3 12V8a2 2 0 0 1 2-2h14a2 2 0 0 1 2 2v4M3 12h18M3 12v4a2 2 0 0 0 2 2h14a2 2 0 0 0 2-2v-4M6 9h.01M6 15h.01M9 9h.01M9 15h.01"/>`),
		g.Group(children),
	)
}