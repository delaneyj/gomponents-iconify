package mdi

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func FormatStrikethrough(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<path fill="currentColor" d="M3 14h18v-2H3m2-8v3h5v3h4V7h5V4m-9 15h4v-3h-4v3Z"/>`),
		g.Group(children),
	)
}