package mdi

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func BeakerPlus(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<path fill="currentColor" d="M18 14h2v3h3v2h-3v3h-2v-3h-3v-2h3v-3M3 3h18v2c-1.1 0-2 .9-2 2v5a6.005 6.005 0 0 0-5.2 9H7c-1.1 0-2-.9-2-2V7c0-1.1-.9-2-2-2V3m4 6v1h3V9H7m0 2v1h3v-1H7m3 5v-1H7v1h3m2-2v-1H7v1h5m0-6V7H7v1h5Z"/>`),
		g.Group(children),
	)
}