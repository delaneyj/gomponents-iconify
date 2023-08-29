package memory

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func Crown(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 22 22"),
		g.Raw(`<path fill="currentColor" d="M2 17h18v2H2v-2M4 6v1h1v1h1V7h1V6h1V5h1V4h1V3h2v1h1v1h1v1h1v1h1v1h1V7h1V6h1V5h1v11H2V5h1v1h1m3 8h11v-4h-3V9h-1V8h-1V7h-1V6h-2v1H9v1H8v1H7v1H4v4h3Z"/>`),
		g.Group(children),
	)
}