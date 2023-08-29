package ri

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func ArrowLeftUpLine(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<path fill="currentColor" d="m9.414 8l8.607 8.607l-1.414 1.414L8 9.414V17H6V6h11v2H9.414Z"/>`),
		g.Group(children),
	)
}