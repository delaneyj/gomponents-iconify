package ri

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func StockFill(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<path fill="currentColor" d="M8.005 5.003h3v9h-3v3h-2v-3h-3v-9h3v-3h2v3Zm10 5h3v9h-3v3h-2v-3h-3v-9h3v-3h2v3Z"/>`),
		g.Group(children),
	)
}