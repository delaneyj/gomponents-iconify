package memory

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func TooltipBelow(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 22 22"),
		g.Raw(`<path fill="currentColor" d="M2 21h18v-1h1V6h-1V5h-5V4h-1V3h-1V2h-1V1h-2v1H9v1H8v1H7v1H2v1H1v14h1v1m1-2V7h5V6h1V5h1V4h2v1h1v1h1v1h5v12H3Z"/>`),
		g.Group(children),
	)
}