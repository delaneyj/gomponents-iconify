package memory

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func ArrowDownRight(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 22 22"),
		g.Raw(`<path fill="currentColor" d="M19 12v2h-1v1h-1v1h-1v1h-1v1h-2v-2h1v-1h1v-1h-4v-1H9v-1H8v-2H7V3h2v6h1v2h2v1h3v-1h-1v-1h-1V8h2v1h1v1h1v1h1v1"/>`),
		g.Group(children),
	)
}