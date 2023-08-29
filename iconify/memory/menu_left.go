package memory

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func MenuLeft(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 22 22"),
		g.Raw(`<path fill="currentColor" d="M15 4v14h-2v-1h-1v-1h-1v-1h-1v-1H9v-1H8v-1H7v-2h1V9h1V8h1V7h1V6h1V5h1V4h2m-2 4h-1v1h-1v1h-1v2h1v1h1v1h1V8Z"/>`),
		g.Group(children),
	)
}