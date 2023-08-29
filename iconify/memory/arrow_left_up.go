package memory

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func ArrowLeftUp(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 22 22"),
		g.Raw(`<path fill="currentColor" d="M10 3H8v1H7v1H6v1H5v1H4v2h2V8h1V7h1v4h1v2h1v1h2v1h7v-2h-6v-1h-2v-2h-1V7h1v1h1v1h2V7h-1V6h-1V5h-1V4h-1"/>`),
		g.Group(children),
	)
}