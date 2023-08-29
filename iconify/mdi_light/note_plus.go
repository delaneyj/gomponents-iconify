package mdi_light

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func NotePlus(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<path fill="currentColor" d="M16 12a3 3 0 0 1-3-3V5H5a2 2 0 0 0-2 2v11a2 2 0 0 0 2 2h13a2 2 0 0 0 2-2v-6h-4m-2-3a2 2 0 0 0 2 2h3.59L14 5.41V9M5 4h9l7 7v7a3 3 0 0 1-3 3H5a3 3 0 0 1-3-3V7a3 3 0 0 1 3-3m2 14v-2H5v-1h2v-2h1v2h2v1H8v2H7Z"/>`),
		g.Group(children),
	)
}