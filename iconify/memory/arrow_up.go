package memory

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func ArrowUp(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 22 22"),
		g.Raw(`<path fill="currentColor" d="M10 5h2v1h1v1h1v1h1v1h1v2h-2v-1h-1V9h-1v9h-2V9H9v1H8v1H6V9h1V8h1V7h1V6h1"/>`),
		g.Group(children),
	)
}