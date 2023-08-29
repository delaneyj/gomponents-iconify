package mdi

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func BeakerCheck(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<path fill="currentColor" d="m17.75 21.16l-2.75-3L16.16 17l1.59 1.59L21.34 15l1.16 1.41l-4.75 4.75M3 3h18v2c-1.1 0-2 .9-2 2v5a6.005 6.005 0 0 0-5.2 9H7c-1.1 0-2-.9-2-2V7c0-1.1-.9-2-2-2V3m4 6v1h3V9H7m0 2v1h3v-1H7m3 5v-1H7v1h3m2-2v-1H7v1h5m0-6V7H7v1h5Z"/>`),
		g.Group(children),
	)
}