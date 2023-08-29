package memory

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func ArrowUpLeft(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 22 22"),
		g.Raw(`<path fill="currentColor" d="M3 10V8h1V7h1V6h1V5h1V4h2v2H8v1H7v1h4v1h2v1h1v2h1v7h-2v-6h-1v-2h-2v-1H7v1h1v1h1v2H7v-1H6v-1H5v-1H4v-1"/>`),
		g.Group(children),
	)
}