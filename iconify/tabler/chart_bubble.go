package tabler

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func ChartBubble(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M3 16a3 3 0 1 0 6 0a3 3 0 1 0-6 0m11 3a2 2 0 1 0 4 0a2 2 0 1 0-4 0M10 7.5a4.5 4.5 0 1 0 9 0a4.5 4.5 0 1 0-9 0"/>`),
		g.Group(children),
	)
}