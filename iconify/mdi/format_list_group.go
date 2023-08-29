package mdi

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func FormatListGroup(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<path fill="currentColor" d="M5 5v14h2v2H3V3h4v2H5m15 2H7v2h13V7m0 4H7v2h13v-2m0 4H7v2h13v-2Z"/>`),
		g.Group(children),
	)
}