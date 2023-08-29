package mdi_light

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func FormatIndentDecrease(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<path fill="currentColor" d="M3 21v-1h17v1H3m8-4v-1h9v1h-9m0-4v-1h9v1h-9m0-4V8h9v1h-9M3 5V4h17v1H3m0 7.5L8.5 18V7L3 12.5m1.41 0L7.5 9.41v6.18L4.41 12.5Z"/>`),
		g.Group(children),
	)
}