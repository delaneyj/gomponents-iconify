package memory

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func ArrowBottomRight(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 22 22"),
		g.Raw(`<path fill="currentColor" d="M17 8v9H8v-2h5v-1h-1v-1h-1v-1h-1v-1H9v-1H8V9H7V8H6V7h1V6h1v1h1v1h1v1h1v1h1v1h1v1h1v1h1V8h2Z"/>`),
		g.Group(children),
	)
}