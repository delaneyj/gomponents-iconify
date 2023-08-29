package mdi_light

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func FormatIndentIncrease(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<path fill="currentColor" d="M3 21v-1h17v1H3m8-4v-1h9v1h-9m0-4v-1h9v1h-9m0-4V8h9v1h-9M3 5V4h17v1H3m5.5 7.5L3 18V7l5.5 5.5m-1.41 0L4 9.41v6.18l3.09-3.09Z"/>`),
		g.Group(children),
	)
}