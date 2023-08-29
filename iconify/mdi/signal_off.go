package mdi

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func SignalOff(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<path fill="currentColor" d="M18 3v13.18l3 3V3h-3M4.28 5L3 6.27L10.73 14H8v7h3v-6.73l2 2V21h3v-1.73L19.73 23L21 21.72L4.28 5M13 9v2.18l3 3V9h-3M3 18v3h3v-3H3Z"/>`),
		g.Group(children),
	)
}