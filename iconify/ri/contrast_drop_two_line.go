package ri

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func ContrastDropTwoLine(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<path fill="currentColor" d="m12 3.097l-4.95 4.95a7 7 0 1 0 9.9 0L12 3.097ZM12 .27l6.364 6.364a9 9 0 1 1-12.728 0L12 .269ZM7 12.997h10a5 5 0 0 1-10 0Z"/>`),
		g.Group(children),
	)
}