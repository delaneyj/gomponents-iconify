package mdi

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func FormatColumns(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<path fill="currentColor" d="M3 3h8v2H3V3m10 0h8v2h-8V3M3 7h8v2H3V7m10 0h8v2h-8V7M3 11h8v2H3v-2m10 0h8v2h-8v-2M3 15h8v2H3v-2m10 0h8v2h-8v-2M3 19h8v2H3v-2m10 0h8v2h-8v-2Z"/>`),
		g.Group(children),
	)
}