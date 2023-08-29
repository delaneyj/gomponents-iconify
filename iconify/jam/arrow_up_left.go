package jam

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func ArrowUpLeft(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<path fill="currentColor" d="M2.757 3.828v5.586a1 1 0 0 1-2 0v-8a.997.997 0 0 1 1-1h8a1 1 0 1 1 0 2H4.172l6.778 6.778a1 1 0 1 1-1.414 1.415L2.757 3.828z"/>`),
		g.Group(children),
	)
}