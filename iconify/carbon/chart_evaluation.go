package carbon

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func ChartEvaluation(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 32 32"),
		g.Raw(`<path fill="currentColor" d="m25.586 8l3-3L30 6.411l-3 3.001zM16 20.5l-5-4.961l1.59-1.57l3.41 3.38L23.41 10L25 11.579L16 20.5z"/><path fill="currentColor" d="M4 28v-2.587L10.414 19L9 17.585l-5 5V2H2v26a2 2 0 0 0 2 2h26v-2Z"/>`),
		g.Group(children),
	)
}