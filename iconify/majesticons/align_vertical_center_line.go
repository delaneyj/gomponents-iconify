package majesticons

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func AlignVerticalCenterLine(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M6 12v6a2 2 0 1 0 4 0v-6m-4 0V6a2 2 0 1 1 4 0v6m-4 0H3m7 0h4m0 0v3a2 2 0 1 0 4 0v-3m-4 0V9a2 2 0 1 1 4 0v3m0 0h3"/>`),
		g.Group(children),
	)
}