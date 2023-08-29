package teenyicons

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func CandleChartSolid(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 15 15"),
		g.Raw(`<path fill="currentColor" d="M1 0H0v15h15v-1H1V0Z"/><path fill="currentColor" d="M8 0v3H7v5h3V3H9V0H8ZM3 4v1H2v5h1v2h1v-2h1V5H4V4H3Zm9 2h1V4h1v2h1v5h-1v2h-1v-2h-1V6Z"/>`),
		g.Group(children),
	)
}