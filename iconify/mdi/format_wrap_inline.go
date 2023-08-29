package mdi

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func FormatWrapInline(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<path fill="currentColor" d="m8 7l5 10H3L8 7M3 3h18v2H3V3m18 12v2h-7v-2h7M3 19h18v2H3v-2Z"/>`),
		g.Group(children),
	)
}