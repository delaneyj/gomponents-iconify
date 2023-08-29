package memory

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func ArrowTopRight(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 22 22"),
		g.Raw(`<path fill="currentColor" d="M8 5h9v9h-2V9h-1v1h-1v1h-1v1h-1v1h-1v1H9v1H8v1H7v-1H6v-1h1v-1h1v-1h1v-1h1v-1h1V9h1V8h1V7H8V5Z"/>`),
		g.Group(children),
	)
}