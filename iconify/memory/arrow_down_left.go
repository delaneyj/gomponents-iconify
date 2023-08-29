package memory

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func ArrowDownLeft(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 22 22"),
		g.Raw(`<path fill="currentColor" d="M3 12v2h1v1h1v1h1v1h1v1h2v-2H8v-1H7v-1h4v-1h2v-1h1v-2h1V3h-2v6h-1v2h-2v1H7v-1h1v-1h1V8H7v1H6v1H5v1H4v1"/>`),
		g.Group(children),
	)
}