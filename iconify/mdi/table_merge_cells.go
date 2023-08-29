package mdi

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func TableMergeCells(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<path fill="currentColor" d="M5 10H3V4h8v2H5v4m14 8h-6v2h8v-6h-2v4M5 18v-4H3v6h8v-2H5M21 4h-8v2h6v4h2V4M8 13v2l3-3l-3-3v2H3v2h5m8-2V9l-3 3l3 3v-2h5v-2h-5Z"/>`),
		g.Group(children),
	)
}