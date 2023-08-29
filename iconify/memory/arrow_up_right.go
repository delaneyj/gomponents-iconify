package memory

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func ArrowUpRight(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 22 22"),
		g.Raw(`<path fill="currentColor" d="M19 10V8h-1V7h-1V6h-1V5h-1V4h-2v2h1v1h1v1h-4v1H9v1H8v2H7v7h2v-6h1v-2h2v-1h3v1h-1v1h-1v2h2v-1h1v-1h1v-1h1v-1"/>`),
		g.Group(children),
	)
}