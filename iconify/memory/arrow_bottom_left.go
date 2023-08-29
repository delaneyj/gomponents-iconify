package memory

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func ArrowBottomLeft(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 22 22"),
		g.Raw(`<path fill="currentColor" d="M14 17H5V8h2v5h1v-1h1v-1h1v-1h1V9h1V8h1V7h1V6h1v1h1v1h-1v1h-1v1h-1v1h-1v1h-1v1h-1v1H9v1h5v2Z"/>`),
		g.Group(children),
	)
}