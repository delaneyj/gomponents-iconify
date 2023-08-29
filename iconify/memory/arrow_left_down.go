package memory

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func ArrowLeftDown(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 22 22"),
		g.Raw(`<path fill="currentColor" d="M10 19H8v-1H7v-1H6v-1H5v-1H4v-2h2v1h1v1h1v-4h1V9h1V8h2V7h7v2h-6v1h-2v2h-1v3h1v-1h1v-1h2v2h-1v1h-1v1h-1v1h-1"/>`),
		g.Group(children),
	)
}