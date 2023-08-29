package carbon

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func ChartMultiLine(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 32 32"),
		g.Raw(`<path fill="currentColor" d="m4.586 28l7.178-5.998l7.994 1.938a2.021 2.021 0 0 0 1.314-.12L28 20.58l-.848-1.812l-6.916 3.229l-7.994-1.938a2.003 2.003 0 0 0-1.74.384L4 25.882V20.49l7.764-6.488l7.994 1.938a2.021 2.021 0 0 0 1.314-.12L28 12.58l-.847-1.812l-6.918 3.229l-7.994-1.938a2.005 2.005 0 0 0-1.74.384L4 17.882V12.49l7.764-6.488l7.994 1.938a2.021 2.021 0 0 0 1.314-.12L28 4.585l-.846-1.812l-6.918 3.224l-7.994-1.938a2.003 2.003 0 0 0-1.74.384L4 9.882V2H2v26a2 2 0 0 0 2 2h26v-2Z"/>`),
		g.Group(children),
	)
}