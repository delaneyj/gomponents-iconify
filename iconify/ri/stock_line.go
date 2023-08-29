package ri

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func StockLine(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<path fill="currentColor" d="M8.005 5.003h3v9h-3v3h-2v-3h-3v-9h3v-3h2v3Zm-3 2v5h4v-5h-4Zm13 3h3v9h-3v3h-2v-3h-3v-9h3v-3h2v3Zm-3 2v5h4v-5h-4Z"/>`),
		g.Group(children),
	)
}