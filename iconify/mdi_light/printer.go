package mdi_light

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func Printer(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<path fill="currentColor" d="M17 4v5h1a3 3 0 0 1 3 3v5h-4v4H6v-4H2v-5a3 3 0 0 1 3-3h1V4h11m1 9a1 1 0 0 1-1-1a1 1 0 0 1 1-1a1 1 0 0 1 1 1a1 1 0 0 1-1 1M3 12v4h3v-2h11v2h3v-4a2 2 0 0 0-2-2H5a2 2 0 0 0-2 2m13 8v-5H7v5h9M7 5v4h9V5H7Z"/>`),
		g.Group(children),
	)
}