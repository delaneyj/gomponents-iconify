package mdi_light

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func AccountAlert(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<path fill="currentColor" d="M10.5 14c4.14 0 7.5 1.57 7.5 3.5V20H3v-2.5c0-1.93 3.36-3.5 7.5-3.5m6.5 3.5c0-1.38-2.91-2.5-6.5-2.5S4 16.12 4 17.5V19h13v-1.5M10.5 5A3.5 3.5 0 0 1 14 8.5a3.5 3.5 0 0 1-3.5 3.5A3.5 3.5 0 0 1 7 8.5A3.5 3.5 0 0 1 10.5 5m0 1A2.5 2.5 0 0 0 8 8.5a2.5 2.5 0 0 0 2.5 2.5A2.5 2.5 0 0 0 13 8.5A2.5 2.5 0 0 0 10.5 6M20 16v-1h1v1h-1m0-3V7h1v6h-1Z"/>`),
		g.Group(children),
	)
}