package ri

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func ArrowLeftDownLine(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<path fill="currentColor" d="m9 13.589l8.607-8.607l1.414 1.415l-8.607 8.606H18v2H7v-11h2v7.586Z"/>`),
		g.Group(children),
	)
}