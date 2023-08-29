package jam

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func PieChartF(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<path fill="currentColor" d="M19.95 11c-.501 5.053-4.765 9-9.95 9c-5.523 0-10-4.477-10-10C0 4.815 3.947.551 9 .05V11h10.95zm0-2H11V.05A10.003 10.003 0 0 1 19.95 9z"/>`),
		g.Group(children),
	)
}