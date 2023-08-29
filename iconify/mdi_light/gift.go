package mdi_light

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func Gift(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<path fill="currentColor" d="M4 13v8h7v-8H4m8 0v8h7v-8h-7m8 0v9H3v-9H2V7h3.04L5 6.5A3.5 3.5 0 0 1 8.5 3c1.27 0 2.39.68 3 1.7c.61-1.02 1.73-1.7 3-1.7A3.5 3.5 0 0 1 18 6.5l-.04.5H21v6h-1M3 8v4h8V8H3m17 4V8h-8v4h8m-3.05-5l.05-.5A2.5 2.5 0 0 0 14.5 4A2.5 2.5 0 0 0 12 6.5V7h4.95M11 7v-.5A2.5 2.5 0 0 0 8.5 4A2.5 2.5 0 0 0 6 6.5l.05.5H11Z"/>`),
		g.Group(children),
	)
}