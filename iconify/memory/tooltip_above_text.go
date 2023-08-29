package memory

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func TooltipAboveText(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 22 22"),
		g.Raw(`<path fill="currentColor" d="M2 1h18v1h1v14h-1v1h-5v1h-1v1h-1v1h-1v1h-2v-1H9v-1H8v-1H7v-1H2v-1H1V2h1V1m1 2v12h5v1h1v1h1v1h2v-1h1v-1h1v-1h5V3H3m2 3h12v2H5V6m0 4h10v2H5v-2Z"/>`),
		g.Group(children),
	)
}