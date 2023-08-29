package ri

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func MapLine(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<path fill="currentColor" d="m2 5l7-3l6 3l6.303-2.701a.5.5 0 0 1 .697.46V19l-7 3l-6-3l-6.303 2.701a.5.5 0 0 1-.697-.46V5Zm12.935 2.204l-6-3L4 6.319v12.648l5.065-2.17l6 3L20 17.68V5.033l-5.065 2.17Z"/>`),
		g.Group(children),
	)
}