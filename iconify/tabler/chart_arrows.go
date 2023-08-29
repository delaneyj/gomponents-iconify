package tabler

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func ChartArrows(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M3 18h14M9 9l3 3l-3 3m5 0l3 3l-3 3M3 3v18m0-9h9m6-9l3 3l-3 3M3 6h18"/>`),
		g.Group(children),
	)
}