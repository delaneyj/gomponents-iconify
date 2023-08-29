package ri

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func LayoutGridLine(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<path fill="currentColor" d="M21 3a1 1 0 0 1 1 1v16a1 1 0 0 1-1 1H3a1 1 0 0 1-1-1V4a1 1 0 0 1 1-1h18ZM11 13H4v6h7v-6Zm9 0h-7v6h7v-6Zm-9-8H4v6h7V5Zm9 0h-7v6h7V5Z"/>`),
		g.Group(children),
	)
}