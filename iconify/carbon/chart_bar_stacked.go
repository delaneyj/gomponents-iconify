package carbon

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func ChartBarStacked(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 32 32"),
		g.Raw(`<path fill="currentColor" d="M4 28v-3h22v-8H4v-4h14V5H4V2H2v26a2 2 0 0 0 2 2h26v-2Zm20-5H14v-4h10Zm-8-12h-6V7h6Z"/>`),
		g.Group(children),
	)
}