package mdi

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func FormatFloatCenter(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<path fill="currentColor" d="M9 7h6v6H9V7M3 3h18v2H3V3m0 12h18v2H3v-2m0 4h14v2H3v-2Z"/>`),
		g.Group(children),
	)
}