package mdi

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func FormatVerticalAlignBottom(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<path fill="currentColor" d="M16 13h-3V3h-2v10H8l4 4l4-4M4 19v2h16v-2H4Z"/>`),
		g.Group(children),
	)
}