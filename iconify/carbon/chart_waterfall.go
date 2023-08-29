package carbon

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func ChartWaterfall(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 32 32"),
		g.Raw(`<path fill="currentColor" d="M28 28V18h-2v10h-4V4h-2v24H10V14H8v14H4V2H2v26a2.002 2.002 0 0 0 2 2h26v-2Z"/><path fill="currentColor" d="M14 4h2v14h-2z"/>`),
		g.Group(children),
	)
}