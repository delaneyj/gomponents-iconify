package ri

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func CarFill(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<path fill="currentColor" d="M19 20H5v1a1 1 0 0 1-1 1H3a1 1 0 0 1-1-1v-9l2.513-6.702A2 2 0 0 1 6.386 4h11.228a2 2 0 0 1 1.873 1.298L22 12v9a1 1 0 0 1-1 1h-1a1 1 0 0 1-1-1v-1ZM4.136 12h15.728l-2.25-6H6.386l-2.25 6ZM6.5 17a1.5 1.5 0 1 0 0-3a1.5 1.5 0 0 0 0 3Zm11 0a1.5 1.5 0 1 0 0-3a1.5 1.5 0 0 0 0 3Z"/>`),
		g.Group(children),
	)
}