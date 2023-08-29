package mdi

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func FormatWrapTopBottom(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<path fill="currentColor" d="m12 7l5 10H7l5-10M3 3h18v2H3V3m0 16h18v2H3v-2Z"/>`),
		g.Group(children),
	)
}