package mdi

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func FormatAlignBottom(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<path fill="currentColor" d="m13 9l2.5-2.5l1.42 1.42L12 12.84L7.08 7.92L8.5 6.5L11 9V3h2v6M3 15h18v2H3v-2m0 4h10v2H3v-2Z"/>`),
		g.Group(children),
	)
}