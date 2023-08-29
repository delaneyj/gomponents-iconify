package carbon

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func ChartWinLoss(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 32 32"),
		g.Raw(`<path fill="currentColor" d="M30 15H16V6a2.002 2.002 0 0 0-2-2H6a2.002 2.002 0 0 0-2 2v9H2v2h14v9a2.002 2.002 0 0 0 2 2h8a2.002 2.002 0 0 0 2-2v-9h2ZM6 6h8v9H6Zm20 20h-8v-9h8Z"/>`),
		g.Group(children),
	)
}